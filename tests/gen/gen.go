package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
)

func main() {
	var err error
	var path string
	if len(os.Args) == 1 {
		path, err = os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
	}
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	path = "/Users/rogee/Develop/myself/atom-project"
	routes := generate(path, true)
	b, _ := json.MarshalIndent(routes, "", "  ")
	fmt.Println(string(b))

	for _, route := range routes {
		filename := "route_" + strcase.ToSnake(route.Name) + ".go"
		routePath := filepath.Join(filepath.Dir(filepath.Dir(route.Path)), "routes", filename)

		if len(route.Actions) == 0 {
			continue
		}

		// 生成路由文件
		f, err := os.OpenFile(routePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			continue
		}

		defer f.Close()

		_, _ = f.WriteString("package routes\n\n")
		_, _ = f.WriteString("import (\n")
		_, _ = f.WriteString("\t. \"github.com/rogeecn/gen\"")
		_, _ = f.WriteString("\n\n\t\"github.com/gin-gonic/gin\"\n")
		_, _ = f.WriteString(")\n\n")

		_, _ = f.WriteString("func route" + route.Name + "(engine *gin.Engine, controller *controller." + route.Name + ") {\n")

		for _, action := range route.Actions {
			_, _ = f.WriteString("\t")
			_, _ = f.WriteString(fmt.Sprintf(`engine.%s(%q, `, action.Method, action.Route))

			if action.HasData {
				_, _ = f.WriteString("Data")
			}
			_, _ = f.WriteString("Func")

			paramsSize := len(action.Params)
			if paramsSize >= 1 {
				_, _ = f.WriteString(fmt.Sprintf("%d", paramsSize))
			}

			paramsStrings := []string{"controller." + action.Name}
			for _, p := range action.Params {
				var paramString string
				switch p.Position {
				case PositionPath:
					switch p.Type {
					case "string":
						paramString = fmt.Sprintf(`String(%q, PathParamError)`, p.Name)
					case "int", "uint", "int32", "uint32", "int64", "uint64":
						paramString = fmt.Sprintf(`Integer[%s](%q, PathParamError)`, p.Type, p.Name)
					}
				case PositionQuery:
					paramString = fmt.Sprintf(`Query(%s, QueryParamError)`, p.Type)
				case PositionBody:
					paramString = fmt.Sprintf(`Body(%s, BodyParamError)`, p.Type)
				case PositionHeader:
					paramString = fmt.Sprintf(`Header(%s, HeaderParamError)`, p.Type)
				}

				paramsStrings = append(paramsStrings, paramString)
			}
			_, _ = f.WriteString("(" + strings.Join(paramsStrings, ", ") + ")")

			_, _ = f.WriteString(")\n")
		}

		_, _ = f.WriteString("}\n")
	}
}

type RoueDefinition struct {
	Path    string
	Name    string
	Actions []ActionDefinition
}

type ParamDefinition struct {
	Name     string
	Type     string
	Position Position
}

type ActionDefinition struct {
	Route   string
	Method  string
	Name    string
	HasData bool
	Params  []ParamDefinition
}
type Position string

const (
	PositionPath   Position = "Path"
	PositionQuery  Position = "Query"
	PositionBody   Position = "Body"
	PositionHeader Position = "Header"
)

func generate(path string, firstStep bool) []RoueDefinition {
	modulePath := filepath.Join(path, "modules")
	if _, err := os.Stat(modulePath); firstStep && os.IsNotExist(err) {
		log.Fatal("modules dir not exist, ", modulePath)
	}

	// routePath := filepath.Join(path, "routes")
	// controllerPath := filepath.Join(path, "controller")

	routes := []RoueDefinition{}
	// travel controller to find all controller objects
	filepath.WalkDir(modulePath, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(filepath.Dir(path), "/controller") {
			return nil
		}
		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		routes = append(routes, astParse(path)...)
		return nil
	})
	return routes
}

func astParse(source string) []RoueDefinition {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, source, nil, parser.ParseComments)
	if err != nil {
		log.Println("ERR: ", err)
		return nil
	}

	routes := make(map[string]RoueDefinition)
	actions := make(map[string][]ActionDefinition)

	// 再去遍历 struct 的方法去
	for _, decl := range node.Decls {
		decl, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}

		// 普通方法不要
		if decl.Recv == nil {
			continue
		}

		recvType := decl.Recv.List[0].Type.(*ast.StarExpr).X.(*ast.Ident).Name
		if _, ok := routes[recvType]; !ok {
			routes[recvType] = RoueDefinition{
				Name:    recvType,
				Path:    source,
				Actions: []ActionDefinition{},
			}
			actions[recvType] = []ActionDefinition{}
		}

		// Doc 中把 @Route 的定义拿出来， Route 格式为 /user/:id [get] 两部分，表示路径和请求方法
		// 通过空格分割，第一部分为路径，第二部分为请求方法
		// 如果没有请求方法，则默认为 get
		// 如果没有路径，则默认为 /{controller}/{action}
		// 如果没有 @Route 注释，则默认为 /{controller}/{action}
		var path, method string
		var err error
		if decl.Doc != nil {
			for _, line := range decl.Doc.List {
				lineText := strings.TrimSpace(line.Text)
				lineText = strings.TrimLeft(lineText, "/ \t")
				if !strings.HasPrefix(lineText, "@Route") {
					continue
				}

				path, method, err = parseRoute(lineText)
				if err != nil {
					log.Fatal(errors.Wrapf(err, "file: %s, action: %s", source, decl.Name.Name))
				}
				break
			}
		}
		if path == "" || method == "" {
			log.Printf("[WARN] failed to get router ,file: %s, action: %s", source, decl.Name.Name)
			continue
		}

		// 拿参数列表去, 忽略 context.Context 参数
		params := []ParamDefinition{}
		for _, param := range decl.Type.Params.List {
			// paramsType, ok := param.Type.(*ast.SelectorExpr)

			position := PositionPath
			var typ string
			switch param.Type.(type) {
			case *ast.StarExpr:
				paramsType := param.Type.(*ast.StarExpr)
				X := paramsType.X.(*ast.SelectorExpr)
				typ = fmt.Sprintf("*%s.%s", X.X.(*ast.Ident).Name, X.Sel.Name)
			case *ast.Ident:
				typ = param.Type.(*ast.Ident).Name
			case *ast.SelectorExpr:
				typ = fmt.Sprintf("%s.%s", param.Type.(*ast.SelectorExpr).X.(*ast.Ident).Name, param.Type.(*ast.SelectorExpr).Sel.Name)
			}

			if strings.HasSuffix(typ, "Context") {
				continue
			}

			if strings.HasSuffix(typ, string(PositionQuery)) {
				position = PositionQuery
			}
			if strings.HasSuffix(typ, string(PositionBody)) {
				position = PositionBody
			}
			if strings.HasSuffix(typ, string(PositionHeader)) {
				position = PositionHeader
			}

			if strings.HasPrefix(typ, "*") {
				typ = strings.Replace(typ, "*", "&", 1) + "{}"
			}

			params = append(params, ParamDefinition{
				Name:     param.Names[0].Name,
				Type:     typ,
				Position: position,
			})
		}

		actions[recvType] = append(actions[recvType], ActionDefinition{
			Route:   path,
			Method:  strings.ToUpper(method),
			Name:    decl.Name.Name,
			HasData: len(decl.Type.Results.List) > 1,
			Params:  params,
		})
	}

	var items []RoueDefinition
	for k, item := range routes {
		a, ok := actions[k]
		if !ok {
			continue
		}
		item.Actions = a
		items = append(items, item)
	}
	return items
}

func parseRoute(line string) (string, string, error) {
	pattern := regexp.MustCompile(`(?mi)@router\s+(.*?)\s+\[(.*?)\]`)
	submatch := pattern.FindStringSubmatch(line)

	if len(submatch) != 3 {
		return "", "", errors.New("invalid route definition")
	}

	return submatch[1], submatch[2], nil
}
