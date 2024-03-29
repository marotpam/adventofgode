package _2021

import "testing"

func TestCalculateMiddleAutocompleteScore(t *testing.T) {
	tests := []struct {
		name     string
		rawInput string
		want     int
	}{
		{
			name: "given example",
			rawInput: `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`,
			want: 288957,
		},
		{
			name: "given input",
			rawInput: `[(<[(({{<{{[{({})}{<{}{}><()<>>}]}[{[({}{})]([{}()](<>{}))}[<<[]()]>[<[]<>>{[]()}]]]}{[[[{[][]}{[]()}]
(<<{((({<(<{([()[]]<[]<>>)}>({[{<>()}]<[<>()]({}())>})){{[(<<><>>{<>()})<{[]{}}>][{[[]()]<[]()>}
{({([({[[([<{(())([][])}<(<>())[[][]]>>{{{{}{}}{<>{}}}({{}<>}<<>()>)}]<<((<><>))<<[]>((){})>>[<{{}{}}
(<[{{{[[[<[[[([]<>)]{{()()}<()()>}][({()<>}(()()))[<[][]>[[]<>]]]]([{[(){}]}(<[]{}>{{}()))]<
{(<<({<{<<[[<[[]{}]>]<[<{}{}>[{}[]]]{[<>{})({})}>]><([[{(){}}{()[]}]])>><{({<[{}()]<{}[]>>(<
<<[[<<[({[<[{{[]()}[()()]}[<{}{}>[<>]]]><[[(<>[])]]{{{{}<>}[{}()]}{<()>[{}[]]}}>](({<((){}
<{[[((<{{[[(([(){}][()<>]))[(<()()>[[]<>]){[(){}][()[]]}]]]<([<[{}{}]([]<>)><<{}<>>{{}{}}>]<({[]{}}[{}{}]
<<<{<[[[[([<([[]<>])[[<>{}][[]<>]]>{{[{}{}]}[[[]()]]}][[<((){})<[]<>>>]<{[(){}][[]<>]}<<[][]](()())>>]){<((
[<{{({<<[{{<<(()())[{}[]]>{[<>{}]<[]<>>}>{[([]<>)[{}<>]](({}<>)[(){}])}}}[[(<{()}{<>()}>)<{<{}{}>([
<{<[{[[<[({{[[{}[]][()()>]((<><>)[[]{}])}{({{}{}}[()])<{<>{}}(<>[])>}}<<(<<>[]>[()])>>)]({{[([<>])(<{}[]>{
(({<{<{[[([<<{{}{}}>(<<>{}><()()>)>]{{(<{}>{{}[]})}(<([][]][<>[]]>{{<>{}}[<><>]})})(<[{{<>()}({}{})}{{{}[]}{<
{{[<<{{(({{{<{<>{}}{[]{}}>[<[]{}>{()()})}}(<({[]()}[()<>])[{<>()}<[]()>]>[[[{}[]][<>{}]]<{<
{{{<<<[(<[({([()()]))[<[()<>]{()()}>((()[])([]()))])][([<{{}()}<[]<>>>](([<>]<<>{}>){({}<>){[]<>
<<({{[([{{<[[<()()>{(){}}]{{()<>}[(){}]}]{[({}{})([]<>)]}>(<([{}()](()[]))({<>()}<<>[]>)>[<[()<>]<
{(<{{([{<{<(<({}{})([]())>)>{[{{{}[]}(<>{})}{([])<{}<>>}]<{<{}()><()[]>}<{[]{}}<[]>>>>}>}<(<<([
([([(<{<[(<[{<<><>>(()<>)}(([]{}))][<(<>{})<{}>>{[()[]](<><>)}]>[[[{[]{}}[[]{}]]]<{[()[]]}[[{}{}]<<><>>]>])
{<<({({{{[({[[<>[]]({}())]{[{}()]}}({[()()]{[]<>}}((()<>)<[]{}>))){[{{{}()}{{}()}}]}]}<<<[[<[]()>][
{{[[[{<<[(([{[<>{}]<<><>>}{{<>}{[]<>}}]<[<<>[]>({}{})]>)({<{()()}>{{()[]}({}())}})}]<([{[<[][]>{()[]}]}]
(<<[<[[[<{{[{{()}[{}[]]}[<{}<>>{[]{}}]]{{[()[]]<(){}>}}}(((<{}{}>{()()}))[<{<>()}{()<>}><[
[[{[<([([<<[[[()()]{()<>}]({<>{}}[<>[]])]>(<([<>()]<()<>>)[[{}()]]>)>])]){({{{[{<({}[]){[]()}><<[
(<[<[(([({{[({<>[]}[<>[]])[([]{})[<>[]]]]([<{}}[{}<>]][(()()){{}()}])}([<({}())[[]<>]>([[]{}]
{([[((<(<([{([<>[]]{()})[(()[])]}])(<[<{{}()}>]>)>)>)){{{<<[<[<(<><>){<>[]}>[[<>{}]]]({<<>{}>
<<(((({<<[{([({}{})[<>()]]<<(){}>({})>)}{<<{[]{}}{<><>}><{<>{}}[()[]]>}<(([][])(<>[]))<[<>[
{{{[{[[([({{{<{}{}>{()}}((()[]]<<>{}>)}}<[{[{}[]]<<>[]>}{[{}[]][()<>]}]<<({}[]){()()}><{<>{}}(<>[]
[[([(([({<[<([<>()][[]()])>{{<[]{}>[<>()]>[([]{})<(){}>]}]({[(()<>){{}()}]}<{(<><>)({}{})}{[[]<>]{{}()}}>
(<{([{{([<{<<<<>[]><[]()>>[(()[])[<>()]]>{({()<>}{(){}})<<{}<>>{<>{}}>}}[[{<{}()>}<(()<>)[()<>]>]{[[<>
<[[{({[<<<{{<[<>{}]({}<>)>(<()[]><[]>)}(<{{}()}<{}{}>>{{()}<[][]>})}[{({[]{}}(()[]))<[[]{}][[]{}]>}<{{<>{
{({[[<{<([<((({}{})[{}{}])){{[()<>]<<>()>}[(<><>){[]}]}><<<<()<>>(()<>)>>[{([]){<>{}}}]>]){({{[{{}
(<<[{[<<({<{<[<>[]][{}<>]>([{}<>]<<>()>)}{(<{}{}>{{}()})<([])>}>{(({[]())(<>[])){[[]<>][<><>]}
[({(((<(<<{({<()()>({}()>}((<>[])[[]]))}>{<(<({}<>)([]())>[{()<>}<<>>])<<([])<[]()>>{(<>{})[<>{}]}>>[[
<<{[(({<<{[{{<<>()>{<>{}}}({[]}{()})}][<[[[]()]]<<{}{}}[<>()]>>]}[<(([{}{}]<()<>>){{()<>}[[]<>]})<{<[
(<[[{{{<<(<<[<()[]>(()[])]<[()()]([][])>><{({}<>)[<><>]}>><[(({}){[][]})((()()))]>){{<[([]<>)[<><>]
[[[[<[<[[[<(([<>{}]<<>()>)(<<>()>{{}{}}))[([()()]([]<>))]>(<[([]())<[]<>>]{<{}{}>{{}{}}}){(<()[]
{[<(<[[<{<{[{({}[])[<>{}]}<[()<>]>][(<{}<>>(()[])){{()()}<()[]>}]}[<[(()[])]({{}{}}(<>{}))>]>}<{{[{([]{})<(){
(<[{{{{{{[{[{(<><>)({}<>)}[(()[]){<>[]}]]}[{([{}<>])}]]}((([{[[]][[]{}]}({<>()}[()[]])][({[
(<<<<[[{{<<((<{}()>[<><>]))<[<<><>>([][])]<{{}<>}([]<>)>}>{{[(<><>){<>{}}]{<<>()>(()())}}[<[<><>]{{}
{({[[<({<{[[[[[]<>](<>{})]((<>{})<[]()>)]([<()[]>({}<>)][[[]{}>[{}()]])]{({<[][]><[]()>}<(<>{})>)((({})
{[([[{<<[{<<[<()()>[{}[]]]>{(({}{})([]<>))[[{}()][()<>]]}>[<(([][])){[{}{}][{}{}]}>{{<(){}>([]<>)}
([{<{{{(<[{[{{[]}[()<>]}(<[]{}>{()<>})][(<[]><{}<>>)<([]{})(<>[])>]}][[{[<[]{}>][[<>[]]{<>[]}]}<((()[])(
<(({{<<[[<<[[(<>[]){[][]}]{({}())(()<>)}]{({{}<>})}><([<<>()><{}[]>][[<>{}]({}())])[(([][])({}<
[<<{{{[[[((<{([][])<<>()>}((<><>)([]{}))>))[<(<[<>()]><([][])<<><>>>)>[[<<{}{}>(<><>)><{[]{}}>}{({
<<{<<[[[(((({{[][]}[{}[]]}(<<>{}><{}<>>))[(([]<>)<()[]>)<[<><>]>]))<<(<<<>{}>>{(<>[])})<[<<
[<<{({{<[<[<<({}[])(<><>)>(([]{})([]<>))>]>]{[(<<{(){}}[[]<>]>>){([<{}()>[()()]](([]())[<><>]))
{({<[<[{(<{[([[]<>]{<>})<{<>[]}([])>]}<{((<>[])[{}[]}){{{}{}}<[]<>>}}<{[[]{}](()[])}<(()())
(({{<{<[[(<[([<>()][[]()])<({}<>)(()())>]{<<<>[]>{(){}}>}>([{({}())}{<[]>(())}]{<<[][]>({}[])>}))]
{{<{{(<[[[{<<<[][]><{}[]>><(())[{}[]]>>}][<(([[][]][<>{}])(<[][]>)){{([][])<{}{}>}[[[]{}]{[
[([{<[[<[{([{<[]{}>(<>{})}]<{[[]()]}[<()<>>]>)}][([[[([])<<>()>]{<{}<>>[()<>]}}((<{}{}>(()<>))(<{}
([<{(<({((<{<{()[]]{[]<>}>(<[][]>(()<>))}{<[()[]](()())>{([][])<<>>}}>))(<[(({[]()}(()[]))([(){}]{<>
([{{[{(<[((<<[[]()](()())>{<{}()}{[]}}>{(({}<>))<<<>{}>[{}()]>})[[((<>)<[]<>>)([()()]{<>{}})]<{[(){}]{[
<[<[<[[[({{{[[()<>]([]())]}}{({[()()]}{{<>[]}<<>{}>})[[{[]()}[{}{}]][{<>()}<{}{}>])}})<<{([{()[]}<[]
<([[<{<<{(<(<{[]}[{}()]>({()}{<>()})){{[[]()](<>[])}(<{}{}>(<><>))}><{{{(){}}<<><>>}}>){(([[()<>]<<><>
{(([(<[<{[[<{(()[])[{}<>]}<<{}()>>>([({}<>){<>{}}]([()[]]{{}()}))]({[[[][]]]})]}>][({[{<<([]<>){()<>}>({<><
([<(((<[<({<[({}{})((){})]{{()<>}<[]<>>}>{[{()[]}<<><>>]({{}<>}[[]<>])}})[[{([[][]](()[]))({()[]})}((((
{<[{[(([<(<{<[<>{}]>[((){})({}())]}{(<(){}>{[]<>})([()<>](()<>})}>[(<{[][]}{<>[]}>[<[]<>><<
[<<{((<(<((<<{[]<>}{[]<>}>{[<><>](()())}>{([<><>])[([]{})(<>[])]))(<<{[][]}{{}()}>[<<>{}>{{}{}}]><[({}{
(<([((({{([[<<{}()><()[]>>[{()<>}{{}{}}]]{<({}{})[[]<>]>{(()[]){<>}}}][{{({}{})}<([]<>)<()[]>>
({([[<<[[[[([<()()>{{}<>}]({{}[]}))]<{(<(){}>[()()])}>]{{<([<><>]{[]{}})<<{}<>>[()<>]>>}{[
<({{(<<((<<(<({}{})<{}()>>{{<>}})([[{}()]{{}()}][<<><>>])><{[<[]()>[()()]]<([]())>}>>[<<<{()}<
[[[(<((({{{{<{[][]}{<>{}}><{[]{}}<()[]>>}<([{}{}]{{}<>})>}<({[[]<>][[]<>]}[[()<>]<(){}>])>}{{(
(<<{({[{({[{(([]<>)({}[])}[<(){}>]}]})}][({(<<{(()())}[<(){}>(()[])]><({<>{}}<()()>){<[][]>[(){}]}>>
([({{{{{([{[{([][])<[][]>}<{()<>}<{}<>>>]}{[{([][])}<<()()>[{}[]]>]{<[[][]]{[]<>}>}}]<[(<<[]<
[([<({(<({[<{[{}<>]}>{[(()<>){<>()}]}](([([])[<>{}]])<[[()()]{[]}](((){}]{[]{}})>)}[<[<{<>[]
[({<[(<<{<[{[[(){}]<<><>>][<(){}>[{}[]]]}(({{}()}<[]<>>){([]()){{}()}})]><<(<<<>[]>{<>}>(({}()){<>()}))[[(<
{{{{<{{<<{[<([()[]][{}[]])[{<>{}}[(){}])><((()[]){()[]})[[{}[]](()<>)]>]}>>}}>{{[[{<(<(({}<>)[<>()]){[
([<{[<<[{({([<(){}>([])][(()<>)<<>()>])}{[([[][]]{{}()})<[{}]({}<>]>](([[]()]<{}<>>){<{}()>[[]()]})})[<
<((<{({[<(<(<<{}<>>{[]()}><<()<>>([][])>){<<[]{}>{<>[]}>{[<>[]][()()]}}>(({{()}{()<>}}({[]<>}(<>{})))
<[(<{{{<[(([{{[]<>}{{}()}}<[<><>]<<>[]>>]({((){})}))<(([<>()]<[][]>)(<()<>>([]<>)))>)]>[[[[[{[
(({{(([{{((<{<<><>><{}<>}}{{{}{}}{[]()}}>[[(<>[])[{}{}]]{<()[]>[[]{}]}])){{{{{()<>}<<><>>}{<[]<>>[[]<>]
(<{(((<<(({{{{[]()}}<<()()>{[]>>}[<{<>{}}{<><>}><{{}()}({}<>)>]}))(([((<(){}>({}()))<{<><>}<[]()>>)<(
<[[<[(<<{((<[<[]()>([]<>)]{(<>[])[<>()]}>[<(()<>){[]()}>{{[]<>}([][])}])[<(([]<>){<>()})([{}()]<[]{}>)>({
<<[[({<<({<{[({}[])<()<>>][[(){}]{[]()}]}[{<(){}>[<>[]]}{(<>[]>({}<>)}]><<([[]<>]({}))<{[]{}}(<><>
<(({([<<(({[{[[]<>]({}())}[{<><>}[[]()]]][(<{}()>)[{{}{}}[{}{}]]]}<{(<()<>>{()<>})<[()[]][{}{}]>}[[(<>[])
[<{<({{(({(<[{(){}})(<<>()><<>()>)>)<{<({}<>)[{}[]]>(<()()>(<><>))}{([<>[]][()<>])[{[]()}{[]}]}>}))<(
[<{<[[(<([<[<<{}<>>>{[<>[]]({}())}>{[(<>{})[(){}]]{[<>[]]}}>][{<{<[]()>{{}{}}}[[[]<>]<(){}>]>([<<><>>[{}{}
<<(({[<<{<([[{[]()}{<>()}]<[[][]]>][(((){}){<>{}})([{}()]<[]<>>)])>}>([[<<(<[]<>><()()>)<[()(
{<({[{<<<{[({<()[]>}[<<>()>[(){}]])<[(<>())<[][]>](<[]()>[{}<>])>]((<(<>{}){<>{}}>{{[]()}([])})[<{{}
[<([(<[{(<([<(<>[])[()()]>{[()<>][<><>]}]<[{{}[])]((<>())[[][]])>)>({(({<><>}({}[]))<{{}<>}([]<>)>)
{<[[({<[([{<{({}()){{}{}>}>{{[()[]]}[(<>[])<{}()>]}}])[[<(([<>]<{}<>>)([{}]{<>}))<({<><>}([]<>))<[<>[]][
{(<[[([[([{{[((){})(()<>)]}}[<(([]<>){{}{}})>{<[[]{}]{()()}>({{}()}[<>])}]][[(<[[]<>]<<><>>>
<{[<{{((<({{{[()[]]([]<>)}({{}<>}<{}()>)}(((()<>)[[]{}])[([]<>)<{}()>])]([[<()[]><<>()>]]<{{[][]}[[]]}>))>)<<
{([([<<((([(<([]<>){[][]}><{[]()}<<>{}>>)(<[<>()]([]<>)>[[[][]]<<>()>])]))(((<{<[]()><()<>>}([[]][()[]])>)<(
{<[(<<((<<({{[{}[]][<>[]]}[{{}<>}((){})]}<([()<>][{}[]])<{<>[]}({}())>>)(<[<(){}>]{{()[]}<()()>}>(<<()[]>
<(<{[({{<<{[({<>()}{<><>})(<{}{}>(<><>))]{<<{}<>>({}{})>}}[({(<>{})({}{})}{<{}{}><{}{}>})]}[<(
{<[([<(([([<(<()[]>(<>[]))[[()][<>()]]>[[{{}{}}(()())][<(){}><{}<>>]]]<((<()>((){}))<{{}{}}
((({{(({{[<[<{<>}<[]{}>>{([]{})<{}<>>}](<{<>{}}<<>()>><{(){}}[[]()]>)>({<([][]}<<>{}>><<{}()><{}{}>>}
[([{[<[[({[(<{<>[]}[{}[]]>(({}[]}{[]<>}))[<{<>{}}{<>()}>(([]<>))]][<(([])<{}()>)[<{}<>>{[]<
([<{(<[<[{([<[()[]]<()<>>>]{[{[]{}}<<>()>]}){[[[<><>][<>[]]]<({}[])<[]()>>](([[][]]<{}()>)[[<>{}]{()[]}])}}{
[[[([{{(({<[[<<>{}>(<>())][((){})]]]{((<()<>>{[][]})((<>())(()<>))){<<<><>>(<>[])>}}}(({<([]
({{[(([<(<([<{(){}}({}())><[{}<>]>]{(<<>{}><(){}>)})>){[<[<(()()}[<>{}]>((<><>)<<>()>)][([{}{}]){(<>[])[[
{([{((({(((([(()<>){[]()}](<{}[]>[{}{}]))<((()()){{}{}})[<[]()>]>)[((<<>[]>[()[]])(<{}<>><<
([[{[[[{<{[{{([]<>){()[]}}}]{({(()<>)({}{})}[<{}()>{[]{}}])<{[<>{}]}>}}>}]]]{<[({(<[<([]()){
((({[<<<{{(([({}<>)<[]<>>][{{}{}}])){([[()<>]]({(){}}[<><>]))((<()()>[{}<>]){{()<>}(()[])})}}<({((
<<<(<{(<([<{[[{}{}]]<{(){}}({}())>}[<<{}()><{}{}>><(()()){()<>}>]>]]>{<[([<<()<>>((){})>][[[()[]]<{}[]>]<(<>{
([{<[[[[[[{(<<()<>>([][]>>)({{()[]}<[][]>}((()()){()[]}))}((([{}<>]([]<>)){<<>()>{[][]}}))]<<[[[[]()][[][]]
{<<[[{{[([<<<{()<>}>[((){}){[]{}}]><({<>[]}[{}])([[][]]({}{}))>>[(([[]<>]){<()<>>[<>{}]}){<<(){}>({}{})>[{
<{(<(<([(<<([<[]<>>(<>{})][([]())<[][]>])([[<>()](()[])])>({(<{}<>>([][]))}[{[[][]]{(){}]}]
{[({{{[([{<((({}()){(){}})((<>[]){<>{}}))<<([][]){{}[]})[{{}{}}{{}[]}]>>}[{[[[<>[]]([]{})]<[{}()]<<><>>>](<((
[[{[(({<[[[({((){})<()<>>}{{<>{}}[()()]}){[([]{})([]<>)][{{}}<<>()>])][{([{}<>]({}())){([]{}){
((<[{(({<{<(((<>)<{}[]>)([<>[]]{()()}))>}>}[[<[{({[]{}}{[][]})}(<{{}[]}(()[])>(({}()){<>[]}))]<<([{}][()[]])>
<(({<[[<{{<<<[{}{}]{(){}}>{<()<>>[<>{}]}>>}[[{[([][])(()())]<[()<>]>}{<<(){}>[()[]]>([[]<>]<[][]>)}][{
<{((<({{({<<{([]<>]{<><>}}((<>[]){<><>})>>{<{[{}{}][<>[]]}>({{<><>}[[]]}<{(){}}[<>{}]>)}})}})<((<{[(<<()(
[(({({{([([{{<[]<>>([][])}<[()<>]>}{(({}())[<>()])}]<({{{}{}}}({{}<>}<[]{}>))>)(<[[<()>((){})]`,
			want: 1685293086,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateMiddleAutocompleteScore(tt.rawInput); got != tt.want {
				t.Errorf("CalculateMiddleAutocompleteScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
