package _2021

import "testing"

func TestCountUniqueDigits(t *testing.T) {
	tests := []struct {
		name     string
		rawInput string
		want     int
	}{
		{
			name:     "given example",
			rawInput: "be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe\nedbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc\nfgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg\nfbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb\naecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea\nfgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb\ndbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe\nbdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef\negadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb\ngcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce",
			want:     26,
		},
		{
			name:     "given input",
			rawInput: "febca cfagb ecbafd efdcbg cbegdfa fg bgafec gfae acgdb gfc | cgf facdeb ecgfdb afcbge\nacbed dcagb cfbega ecgdab be deafc bfgdca abe edafgbc debg | dceab dgbca gedabc fecabgd\ndecfb bgadec dgefa bcgfade cae adecf ca eacfdb bcaf dfgbec | bfca fdacbge cedfb dfebac\negd agdcfe gedbfa gbcfae ed dafgebc bdcgf gdbfe fgabe dbae | cefbga egfcba dbfeg ed\nbfeca beadgc gdfeacb aebcfd eg aefg cgefab ebgfc gbdfc beg | bedacf eg dabegc bfeac\nfcd fd fgdceba edgac cegdbf gcafdb agdfc cafbg beagfc dfab | dacbfg gfacb fcd fd\ncfbgeda gdfae abdfg dfceg edagfb geba bfcaed fcdbag ae aef | dcbefa dagfb dcefg abgdef\nedafbg ebg gbdea cdbefga bfgcea ge egdf aedcb gfadb gbcfda | degbaf dbgafe ecdab dagbf\nadecfb aebdc ab efbagd cfbed cagde baecgfd egdbfc bad bcaf | cbaf beadfgc febcd gaecd\ngfdacb bgdace gb cdfbg fcbde fdgeac bgfa dgb acfgd bdcefag | bg dbegac fbdegac agfb\nbfdace bdcfa ebcag egcdbfa dbg dg fbegdc bcgda fgabdc gfad | cadgb abcdg eagbc dagf\nbed faegcb fedcba dgbcea feabc cdefb bd bcdfgea edgcf afdb | dgcfe adfebgc ebgcad cfebd\ncfead dc dagef gfdc cgdefa acdbeg cad egfabcd bgeafd aecfb | dafge fdgeba faecd cbaef\nafgbde gebca ecbdaf caedb faceg abgcedf cdbeag cbdg bg agb | bg bcaed aecgb cgbae\ndfab dfgae bcafge cbefagd fdbeg fa fea fcbdge aegcd bdeagf | gfaedb af acbgfe faecgb\ndecafg cadef fadecb dfeba aecbgd gefbd aeb ba ebfgdca abfc | gbfaedc bdefg acbf gcafde\nefgbd efdacg cd gefcd becagf dgc edca egdbafc cegfa fgdbca | eacgf bgafcd bfaegc dc\ngabed gdacfe egabcdf cdaf fdg acgfe fgdea gafbce degbcf df | fcaeg cefag bfegdc dfg\ngd badgcf cdg gbdcae cbdea dgfceba bfadec gdbe geadc acefg | gdc ebdg dg aegcf\ngbae fedcag bg ecgbda bfecd dcage fecgbda bgdce acfdgb gbd | cbgafd edcfb dbg bedcg\nabcgdf cfgae cebag gedbafc dcgbae abg decfbg gcebd ab dabe | gdbeac ab ab dcfbga\ngab efgbd gbcfad edgfac ba gaced dgcbae edbag fcgedab aecb | gfdace adebg ba cgbdfa\nacebg aedbfg deca ad dcaebg cgfbd abd cdabg dbecgfa agfceb | ad fadegb dcae cbfgd\nfdagcbe ebgdac ecfgbd fbc fc cbged fbegc cdef cfdgba agbfe | bfgce begcf cdfe cbf\ncbd fadgb dceag cb gfadeb begfcd cafb cgfdba gbadc cegdbaf | bdc bc adcbfeg dgace\nbgdefca cbgadf ebacf dfbac gedcfa cef fe gcbae fbed fcdbae | fe cadbfe ebfgacd cdbafg\nacfde cgeabd cag gfcb adgfc fdgbae fbdag cgefabd dgbacf cg | fgadc dcbgae ebdagf gbacfed\nebdcg cfdbage cefbag cb ecb dabc cagdeb gbefd afdcge adegc | dgbafec dcfgea dgbef aecdg\nacd bfcged aedgfcb afdbce efcdb cfdae cbae cdafgb gdefa ac | aceb ecbgfad bdaefc efbcd\ncebd aefdgcb bafecg bfagd adfecg dcgab gcd cd ebcadg cebag | cdg cdg gcdba dafgec\naecgbf bgfae gbcfeda fba abce gebfcd ab gcebf gcdfab gdafe | begcf baf baf fbgcae\nacdgfb bgedfa cbadfe degf gbeac gd adg dbgcaef abged febda | dg dg ebdfca fgacbd\ncfb fbdgac gcdaefb dabec fbcda fagbd cf cdgf bacfge ebfagd | bcdae bgdeafc acegbdf cf\ncaegdb acgbf abdcgf fgebd afcd cd ecafgbd gcbdf dgc fbgace | acfbg efacdbg bcfgd fbcgad\neab fecbdg dcfab fdcgaeb gdebc cegadb bcaed bfgace ea aegd | ea bfadc bcgdae becfag\ncg fdbeg eafcb bcefda cabfdeg cbfgae bfecg bagc gec dfgaec | bagc ecg cg abcg\nceafg dgeab ecd fcbedg edcgfa agdbfce dafc bcfage agcde cd | ecdga gecad cdaegf dafc\ngedcf bcad abcfgd bd dcfbg bfgac fecgab bfd efgdcba faedbg | bd bd dacb cfgbd\ndcgfab fabed decbf gebcdf da efbag deca acfbegd dba fadcbe | ad abgef efdba debfc\ngadcfe gefdbc faced cfa geac afcgbd efbacgd ac cgdfe baedf | ca afc fcegd gedfca\ngefca ecdagb bg ebadf dbfg eagbf fbegdca cefbda beg fbedga | gb fabge gaebf ecbdag\nabd gdecbf da ecfbd afgbde ceda abdcf dfbcgae dbacef facbg | dafcb bda dgaebf gfabc\nafbegc cdgea aedfbg cgdb bgcaed egcdbfa cfeda cg dagbe gec | agedfb fabdge fegbda ecg\nac acdbge dgbefa abc dbecfa cabgefd fegbc fbace eadfb fdac | ecbaf ca cgebadf afcd\nfdgcb cebf dcfbge cbdag gdbef acgdfe adgbecf fgc badfeg fc | ebcdgf cfg bgdac gfdeab\ngefc ecadgb efcabdg fca ecgad fc ceadgf bdfacg afebd cdaef | fca bgdeca efbad cf\ngcbfd ecdbf cfdgab gacfb dg gbda efcgba gfd egdcaf ebcadfg | gd gbacf facgb fcbed\necfa edfgcb bgdae abcdef ac debfc cad acbed febdgca bfcgad | edabc ca acd eadbc\ndbceaf cbagfe bdcgae bdc cdag debgf abcge bdgcaef dc cdebg | bdface gadc bgaec cebgaf\necf dfbae aecbf gabcf bfgdca cbegfd ec gcea cebfag cedabfg | afdeb efbca dbfea afbcdg\nbdagcfe gd debgca dgfc gbaef adbfce fcdbga gad dabcf fbagd | gdbeca dga cdbafg deabfc\negabfc fdeb dcbae cfdea aeb be dabgc abcdef cagedf fegbacd | bdfe cedfag edbf gaefcd\ndecfgba efgdc befdc dcabf ebc gbacdf acgbfe abedfc be deba | fdebc bec eabd fdagceb\nefbag gdae dgcbaef cdfbge gbfed gdcabf fedgba bag ag acebf | gbedf bafdcg cabdfg dgcfaeb\ndcfb agcdb aedfcg gfbdae fbgad dac cbeag dfbecag dc fdacgb | cagbe gadfbe dcfbga bgcad\ngefcb edcfg bega befdac bagfec cadbegf bce bcfag dcgafb eb | gbceaf bgcfa acefgb abfcg\nbgfce agbfc fga dabceg bdagfec afdcgb gcdba adfc fa baefgd | gfacbd fa agf gacbde\ngbfa bef bgeda cbeadg begdf gfcdaeb egcdf dfcaeb bdfage fb | fb dfbeca gfecd abgf\ndbfeac ecbfg ec dceg afbdcg gefbcd dbcfg geafb ecb bdegfac | fagdebc cadgbf gbdfc ec\nfgdbe beadfgc dfeab gbaf gf gdefca degcb aefbdg facbed gfd | fdcgae dgfbe fgd bfdae\nbaecfd aeb cfae ae dgbaefc facdbg fbgead gbced bfdca bdeac | afce caef ecbda ea\naedgb agf abefgcd fabgd afgebc faed dgfaeb bgfcd edbcag af | af adefgb adfe abdgec\ngdaecb gbcad gbcae cgfabd ebgfc ea bagcfed gae dbfage cdea | abdcg gadbc eacd adbegc\nfcdbe febacg gcfd egf ebgdf fecgdba cfdeba gf agbde fbcdge | gfe abged efg cdfg\ngb gdceafb bceaf cebgfa aedbcf fgb acgfb fcdag eagb cfbdge | efagbcd agdcbfe fbdegc abge\nebfac cbeafg df dcbag acfdbe cfdbeg dfb gecfadb feda fdbac | gdcba gedfbc gedabcf cagbd\ncbfdg gbacf gdab bcdfeag dbf fgdec cfdabe dgbfac fgaceb bd | dcefba fgdce edfcg bafcg\nec dacfeg aebc dfaegb eagdb gefcabd aedgcb dbcfg dec egbdc | fbadecg dfecag efgcdba dfbaeg\ngafbdc cbafde cfgebd gabcd cebag ad cad cafdebg bcgdf dgaf | fagd efcdbg acd bfdceg\ncde afebcdg bagcd ce fagedb dcfgea afbced efcg cgeda gfdae | acedbgf cabfed cefg dce\ncabg geabf gfbeca dfgea gfbeacd dgfceb gbe dcbeaf befac gb | agdfe dcfeab adefg eafgd\ndebgacf edgab agedcb fdega ab cadfbg bad cbea bgecd dfbcge | adbge dabegc fagde abd\nbdefac afgbc abcef ce bdefga dfbea cef bgdfaec ecdb agfced | dcbe caebfd dgafec acbgf\naecfbg gdaefc becgadf bdgfa cg efbcad decaf fcg acgfd edgc | cfbdgea cg cefdga cfg\nfd dgfcaeb fedg dcgbe bdgecf bfcde abcfdg dbf becaf bgeadc | cdgbea df egdcfab cdbfe\naedcfb bdecg bcdefg cd fbadcge edc gdcf gbdef bfagde begac | befadc bfeadg fgdbe dc\nde cgadf gde gecfb egfbac bfde dcabeg ecbdgf efdacgb dcgfe | defb bgefc ged cbagfe\negdfba cdafbge gfa dfgab cbdaf becgad bcgafe gfde bagde fg | bfcedga bacgde geabd cfbda\nbacgfd gec cbfeadg gdeaf afceg ec dagecb gacbf egafbc cefb | abgfcd fdabcge egc abcgfe\nedbgac fecbg adcgef fdbgc cbaeg feab bceafdg cgbfea fe efg | abfe fdbgc aebf cgadfe\nafbdcg dacbf cdbgf acdfe ebafcg ab beadfcg cfgbed cab dbag | becfgd abcdf bdcfaeg egbadcf\ndafceb gebfad cebg gbdfeac dfcga dgbac acedb gbd ebgcda bg | deabcg bcgade egcabd bdcafe\ngfa agcfe fbcae cedbaf cdfge ga gcfaedb afbedg cgba bgaefc | bagcfe ga fbace fceba\neg cbaefgd dagfb adcbeg aeg cefabd bgec dgeafc abdce dbaeg | fcagde bdcage gabfd acdegb\ndfeagcb ebf eagfc gbefac cgdfea dbfag eb dcbafe gecb egfab | eb dgafb faecg eafdgc\nbefdc gfbceda ecdgba bfdacg cde gecfb deaf bedafc ed cbdaf | becfd bdacgef eacgdfb gbdace\nagdb cdafegb cfabd cedbgf cdafeb fbgac bgc gb adcgbf faegc | cbgadf dabg fgeca cbgfa\ncfeabg cfeab fbge gcdae fgc aefdbgc gf gaefc dcgfba ebdfca | abefc cgf fbdaecg bafce\ncbfgda ebfdcg dgbfa bd fgdca abecgdf dgb cegfda bgfae cbda | gdcaf fdgab cbfaedg cgaefdb\nef cbfgae abcdf adecfbg gacbe abecf agdbec ebgf dfceag cef | efc ecf dcbaf cbaeg\ndgcaf bgfd aebcd ebacgf cdafb fdabecg bdafgc adecfg bf fcb | afcdb fcadg fcadg adgefc\nafcgeb fdgc bcaefdg egacf ged gd faebdg cdaeg badce cfaged | cedba egcbaf dfbagce eagfc\nedgaf aegdfc dgcfba befdag cd cafdbeg cged facde eafcb fdc | gebfad gefda fcd agdfe\ndebfa dcgbefa gdcbae cdfbae ab bdcegf cebfd bea cfba agedf | fgecdb bcaedg feagcdb fecbd\nabfecg fdegca dag cead dbfge gaedf da afcge cdgfba gbdcafe | gfedb bgfdac dag agdfe\nacedgb dagcebf abgcf acgbd fbgd afb cgadfb ebfcda fgcae bf | agcfe fgeac ebfdac bdcafe\nadecfg cbda ecabf egcbdaf fgaedb gcbef cfbaed bfa ab efdca | fbcae fcaeb dbac cfadeb\naecgdf bgcea dbc facdgeb efdb db edgcb cdgfe cgdfba gcdfbe | fdgec bd gcdabf dbecg\ncbgdfe gbdfa dbe bfgdeca cgbdae cbefag becag ed aedc agebd | ecgab fgbdce aecd dace\ncfbed gfcbea abdefc gbfcd eabd dce de geadfc cbgdaef febac | ed egcafd cde edc\nfgcde fcbad fdegca ae fdcea dcegbf geaf eac badecg aecgbfd | egdcfa dbfac gafe ae\nbgadcf acgdf afgedc bdgeacf bfd cbaed fcgb agfdeb fb adbcf | fdbgca fdbcag acdegf cagfde\ngacbfde ebgdcf eabcfd fae fa bfead cabf cbefd gfcaed gaedb | fa fcdeb fecagdb bfac\nadecf cafbd bcdagf gbadc bf dbf gdebca edfagbc bfga dgcebf | fadbc dgcfbe dfb feacd\ndfgae fgcedba ge bcfade caedfg dfbga fdcea gaedcb gea gefc | cfeda fbagd gea eag\ndafceg dfba cdeabf da bedgafc adecb gecab fdecb febdgc cad | caebg cbfde abdf ecgab\ndaebg cebg bcdae gb bdg cdfgba bfcagde dbaceg fdecba fdega | geadf dabge gebda ecabgd\nafedbg feb bf acfb dbeca bfegcad cdefg ecfdba dcfbe bedgca | facbedg cfegd cfba ecfdg\nbd abgcf cfgebd cfgdae abde dafceb fdb daecfgb fcdab cdaef | bd cbfda cabgf daeb\nge cdfga edga caedfbg edbgfc dagfcb efcga cfagde feg faebc | gef eadg fcbagde efgbcd\ncabegd bec eafbd bagc dfcaegb dcgae cb ecdba bdgfec egfcda | cegdfb dgcae decfgb bcdea\nfcbgead af cdbga fabegc cbafd fead cefbd fbcdae gdebfc baf | cadbf cfdab efbdc baedcf\nedfbag bgc cfba ebgcaf cegdfab gfbae eadcgb fecgb fedgc cb | bdgeca cb cb dfgce\nbfadecg fad baedfc cabef gacfbe agecdf cgdfb aedb da acfdb | acebfgd aecfb efbca agbfdce\ndefbca ebcdgaf fgbe gdcfa degcba decbgf fe defcg ecf bdegc | eadbgc fgbe gedfc dcegb\nbagecf bga gbde gb dacbf gefacdb edcagb cfgeda gdace cadgb | dgacb bg bdeg acdeg\nbdc dgbafe cabdeg bfgde befac defbc cd fcdg afgecbd gbfced | cbadeg cdfebg gaebcd gefdb\nafgce dc fabged gbafd fgcda gecdbf badc fdcebga cdf adcfgb | cfbdge dabc dcagbf dgfabe\ncdafeb egf afbeg dagbef cfbgead edag cgebdf gbcfa daefb eg | dabfe fbacg efg eagbf\nedbf afcbedg eacbgf fd efadc fagbdc abfcde daf becaf acdeg | cgfdba efabc adf fgcbda\ngcabe bcadeg dag dgafcbe agdcb faegcb dcbaf afebdg egcd dg | gdfeab cagbd bfceadg gda\ngafcd cbgafd gbafec fcgade bcfgeda cdea bedfg ea fae eagdf | acgdfe deacgf efgdb fdgac\nfecagb ae cgfbe gdafb fecdbg faegb gbcefda aefc aeg gecbda | gae ebgfc edgcfb adbegc\nfgced gcdfbae fec edcb ce bgdfe dfagbe acgdf cbaefg gbcdef | cdafg gaefcb becd cef\nbegdc bacegfd cadfeg afbgc cbdagf egbcf feg bgacfe feba fe | fcbeg gbcaf egf fe\nbed eagdb fagdb efadbc aefbgc ecgd ed gbcdea cadgefb bgaec | bde de ed ecdg\nedbgfc ebfacg cdgef ga daegcf gfa bfacd aegd fgadc befgcad | cdbfa cfedabg dagcf gfbaecd\ngdacb bcdgfa egacdb facb cf cfd cgdefa gbfde egcafbd bcgdf | cfd aegcdb fdc cbfa\nfdca cagfb ac gcdbef bcfgade ebafg bdgcea abc cgfabd cgdbf | fgcbd febdagc bdgaec gbedfca\nca dgcfb fgead gca edcgfa fagcd gecfba aedc dfeagb acegdbf | ac cga cade ca\ndef gfbdc ebdgaf cbde gdcfe eabcdgf ecgfbd dbfgca gfcea de | dbec fecga dfbgc dcfeg\ndacfb bed gcbaefd fbcdag afdecb befc edfab agefd dabecg eb | dbe bdgfac abedcf cdbgfa\nabf cbadf cfagdb acgf fa fdbcg bfagde befcdg dcegabf cabde | agcf fa bdaec cfagbde\nedbcag aebcf afgb fcgbde beagcf bea cgefb ba edgfacb faced | aeb gbcdae ecbfa decbgf\nad adgfcb afgdc ebafcg eadgfcb facgb ecgfd acdb gda gdeabf | fcegd cdgfa abefcg dcafg\ncfebd gfdec feadg gec gfabed cefdga cg gdca gebcfda aecbfg | cbfeag dcag gc cabdfeg\neafdc fac af ebdcf bagcfd gfae eagdc bcfegad gebdca gcdefa | fcbgda dfebcga gdbace defcb\ndagcf aecdgbf bedfgc bafedg adcbg gdcbea cgb dabge cb ebac | fgdebc edfabg fdcag ebac\nbgdaf cbdgef bdcag ac abgedc afgceb cdegb acde ecgbfad cga | adec gbcda dbcge edac\ngbafe gabecf gaed edbgf ed abedgf bdgcf baefdc def cbgefad | febga agde fed gfdcb\ngefcba fgb gf eabdf gcabd edfacbg bcfadg dgfc dfgab gdeacb | fgadb gcdabf fg gfb\ngbadce cbdafeg ce egcdfa agdeb fbdac ecgb ebgdaf ced acedb | bcafd ecbg bcadf gedbacf\nag fbcgead baefcg cfbgd feadc adge fgacd dcagfe dcafeb fga | gaed abdfce ag ecfbga\nefbgdc cbagf gcdabef fegcda gcdaf daeg gd deacf fgd acedfb | fadgc fgabc befdac dg\ned gcefa dcefga fcdbgae dcgea efbdac acgfbe dbacg aed gfde | ed efcgba adgce de\nafceb adbfge beafg bga bedfgc ag bcadgef dgbfe gdaf geabdc | ebgcdf gfcbde ga edbfg\ngafcde agcfeb fdcge ebdfc gebfcda gef gf fgda eagcdb edcag | fabceg bgcfea fagd ebcdf\necadbf bedfg da dgfbcea bacfge ebfad bcaef deagbc adfc ead | dae fadc fdgceba debgf\nadb fbcae ad acfebd gefcbad cefgab bcdaf gdfcb ebagfd eacd | dbfca fdegbca abfecd cdbfg\nfbdcga ebfcd fbadc ca fabgdec bca bagfd bcgeda agfc bgefad | gafbcd cbedf afbdg afbdgc\necfdba cd facegb gbacde gcbed dce cegba bcdegfa befgd cgad | ced cd fcagebd dce\nebdca fcabge dceagbf dgef fbedc ef feb cdfgb cdbfge cgafbd | bcade gfbcae cfgdeb efb\neda egabfd egfacb da fdbegac gdecb adfc dacegf acdeg feacg | dcage cgbde faceg gcefa\nefcgda fgec dafebg agfcd aefgdcb gf ecadbg cgead fga adfbc | fegcda afdebg dbafc fag\ngbdaec gaecfd cbdag acfbd cfegabd fd fcd cgfadb fdbg fbcea | fd ecbagd fdgbeca ebdcfag\ndbgacf efgad abcf dcebagf ebgacd cfbdg dgcefb ab dfbga abd | dbgeacf cfdgba gbdecf gbfad\nabegfd cdafb bdfag ecagdf ag acegfdb ebdfg fag cdbfge bgea | fga dcegfa dgfeb fceagbd\ncd acbefg cdfge dfaceb gcaefbd cfd fcaeg egdbf dacg daegfc | egcfba cadg cd dcf\ngebdfca cdbefg afcg gecbf ag debacg dbaef fabge aeg gbcfea | cgbeda gfecdb ga cgebad\ncgdfeb dfacbg ecd cadfg cadbgfe aegdb fedgca acef cadeg ec | cgdae agbde dfcebg cgfdeba\nfg bdcagf ebacfg gbacd cgfd gfa dgbfa adcbeg edbfa dfabegc | aefbcg agbdc fdcg agcdfeb\nba cbgefd fgace bae cadb edbfc abfedg caefdb caebf gaebcfd | bae cadgefb efdbc ba\nedcfb cdabe gacdeb dabfecg gefacd aedgbf dac bgac ca edbag | facedg dfeagc daegcb gdabce\nafbdce gfc fecagdb egdaf gfbaec bdcfe bcdegf efgdc cg bcgd | gbcfea dfgebc cdefb cg\ncgdaeb dfbacge cgdbe acb gcdfeb dfcga ba fcadbe gbae bcdag | fadbec cadgf ba ba\ncgdeb ebfcag gbdafc bgfde gecdba ec abgdc bce dgecafb cade | cdgafb dcea gdabfc bcagdf\nfadbec gecfbd deg ge afdgce bdfec agbcd ebfg cbedg fbedcag | bgef abcdef cgbde cbefgd\naegdb befgcd gec ecgba aecd cfbag bdefga ce gacdeb dgecabf | eadc gbfced bgcae dgaeb\ncfabd fbedac cadfgb eb bce cdegf decgab eafcgdb cfbde afeb | fdbec fbacd be ebcadfg\nbdecga gdac cdbefg dc bcd baecg bdeac efabd fgceab dbacfeg | bedac faebd baced egbdca\ncg cafgdb bacfde gfaeb cfadb bagfc fcg ecbagfd fdcega dbcg | febag fcg bgcfad gadbcf\nfbde fgd df bagef gacdebf cefabg agdec gdafe gaebdf bcdgaf | efdb abfeg febd debafg\naebcgd abcdfe bdfeg eaf dagecf cgaf bfacdeg fa gdeaf gedca | cedag gedfb agbecd afcg\nfedbcag cgfabd gba faegd eadbgc gceb bgeda bdcea gb aedbcf | beacd ecbg bdaec adgbec\ncfge gc gebfcda fbadc fdbcg degfb gebdac gdeafb dcg fcegdb | dcbgea debgf cdbaf bgcfd\nceafb fegbdc bcfed ecd daefbg dfgecab cgfd dc aedgbc begdf | edbfg efcab cdfg bdgef\necag fegcba bfgac cdbfg efabc befdca afg gcbaedf abefgd ag | fedcba fgcbd ga acge\nbadgf cgaeb ebdga afcbde ebdgaf fged acfgdbe de ebd fdbacg | gfdbea badefg dgeba beagc\nbeagcd cga cafgedb ac febga fgcba cgfbd dgfcab gcefdb cfda | ac fcda agc ac\ndabefc abgfdce decfb gf gfed cbdgf bdegfc gfc gafcbe cagbd | cfg dcfeb bcefd edfcb\ngedab acd ceba agdcb bagdce fcgdb ca dgaebfc dfgbea fecagd | adc acd adegfc cda\ndg eafdg gcbeaf dcfaegb acgebd fadgbe dgbf daecf fgabe deg | gdfae bgfae fgbdea becdag\ndecgfb bcgfdae agcbd afcdb fb ebaf fbd cedabf ecdfa ecfdga | dfeac fbae fb degabcf\nfeabgdc gfdbe bgcad cf geacfd dafbcg cbgaed fabc cfg cdgbf | dbcag cf cebagdf gfdbe\ngdbef ecdgb efcd cdgba fabdeg fdcbeg faegcb ecg fdegcab ce | bgefcd gcdba bgdef egc\ngfebd abd cdfbga gcdea edcbfg abfe ba gebda dbeacgf eafdbg | beagd egdba ebdfg fdageb\ngafbcd bc bac fgeba fbdega agbec febc egadc bcfgaed beacfg | agbfe cfbeag gceda adegbf\ncdbg gfcead db gcbfde bdefc gefdc bdf fcdbega fecba dagbfe | dcabgef fegcd bgdc cgfbead\nbagcde efbcgd adgef fgaecdb bfeacg eabfg cebfg eab ab cfab | eafgb bacf ab eba\ned dcfbeag dfcga cedb cbfae fgdabe dbcefa fbagce dfe dacef | afbgec de cadfg de\ndgc abfedg bdgea cbfde dfcgab ecag gbcdae cg gdfbcae gcbde | afbdeg dgabec begad gdc\nbefcd dg bgacf fbcegad dgb gcad gbcdf fagdbe bfcage gfdcab | gdabef fdecb gcda agcd\ncb bgdcfe bgacefd fdgbe edcgb cegad cbagfd dbefag bdc efbc | cgade gbfaecd bcd cbdge\nadfcge afgd acfeb efcda da gdefc cegfbd bcegad fdbcage cad | defgbc aebgdc dgfa cadfeg\ngacdf bdg aegbdc bg debaf bfeg gadfbce gfbdea cfbeda gdfba | bg fgdca bg gb\nadbecg dfa fd gafbe gefda cefgda ecdag efdc abdfcg egbcdaf | fegba gfadec fadeg efcd\necfgad gbae ecdag ba agbdec bac fdabgc edabc cdgeabf cbfed | dcfeag cbegad cefadg gdebca\nfgbca cge gcefb gadebcf eabcfg gcadbf feag eg fecbd egdabc | ge afgdbc fgdbac egc\neadg eadbc gfbcda badefc gdc fcgbe dcbge acfgbde dacgbe dg | cgd egbcd gcdfba agdbce\nga befgac fcgedba fabced agfce edgfc gbea aebcf acg cagfdb | gecdf afebcd bgdfac cga",
			want:     301,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountUniqueDigits(tt.rawInput); got != tt.want {
				t.Errorf("CountUniqueDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}