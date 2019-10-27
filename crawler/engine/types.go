package engine

/*
 *
 * <pre>
 *  Version         Date            Author          Description
 * ---------------------------------------------------------------------------------------
 *  1.0.0           2019/10/27     redli        -
 * </pre>
 * @author redli
 * @version 1.0.0 2019/10/27 10:44 AM
 * @date 2019/10/27 10:44 AM
 * @since 1.0.0
 */
type Request struct {
	Url        string
	ParserFunc func([]byte) ParserResult
}

type ParserResult struct {
	Request []Request
	Items   []interface{}
}

func NilParser([]byte) ParserResult {
	return ParserResult{}
}
