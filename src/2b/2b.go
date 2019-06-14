package main

import (
	"fmt"
	"strings"
)

func checksum(input string) {

	var splitInput = strings.Split(input, "\n")

	for i := 0; i < len(splitInput); i++ {
		// iterate through remainder of list
		// see if there is one diff between each string
		for j := len(splitInput) - 1; j > i; j-- {
			var numDiffs = 0
			for k := 0; k < len(splitInput[j]); k++ {
				if splitInput[i][k] != splitInput[j][k] {
					numDiffs++
				}
			}
			if numDiffs == 1 {
				fmt.Println(findCommonLetters(splitInput[i], splitInput[j]))
			}
			numDiffs = 0
		}
	}

}

func findCommonLetters(firstString string, secondString string) string {
	var resString = ""

	for i := 0; i < len(firstString); i++ {
		if firstString[i] == secondString[i] {
			resString = resString + string(firstString[i])
		}
	}
	return resString
}

func main() {
	checksum(input)
}

const testInput = `abcde
fghij
klmno
pqrst
fguij
axcye
wvxyz`

const input = `tqzvwnogbarflkpcbdewsmjhxi
tqyvunogzarfckpcbdewsmzhci
tqyvunojzarflkpcbdcwsmyhxi
tqyvunogzarclkpcbdewmmjrxi
hqyvunogzarflkpcbczwsmjhxi
tqyvunogzarflppcudewsmjhei
tqrvunogzarflkpcbpewsmjhdi
aqyvunogzarflkpcbdewsjjsxi
tqyvtnogzarflkkcbdewymjhxi
tcyeunogzarfkkpcbdewsmjhxi
tqyvunozzarfvkpcbdewsqjhxi
tkyvuwjgzarflkpcbdewsmjhxi
tqevunogzarflkpnbdkwsmjhxi
tqyvunogqarflkpcppewsmjhxi
tqyvunsgzarflkpcbrewsmjhxk
tqyvunogzarffkpdbnewsmjhxi
tqyvunogvarflkpjbdewsojhxi
tqyvgnogzarflkpfbdswsmjhxi
tqyvunogzarfxkpcbtersmjhxi
tqyvukhgzarflupcbdewsmjhxi
tqyvdnoozyrflkpcbdewsmjhxi
tqyvunogzorflkpcbdewsvjhxd
tqyvunzqzarflkpcbdewxmjhxi
tqykunogzarulkpcbdhwsmjhxi
tqycdnogzarflkpcbdewsmjhxz
eqyvunogzarflkpcbdhwqmjhxi
cqyvunogzaralkpcbdewsmjvxi
vqyvunogzarflklcbeewsmjhxi
tqyvunogzarffkpcqdewlmjhxi
eqyvunogzarflkpcbdejsmahxi
tqyvunogjarflkocsdewsmjhxi
tqzvunogzanflkpcbdewsmjhxg
tqycunogzarflkpabdewsmkhxi
tqyvunogzarnlkpcbdewimjaxi
tqyvunogzarfzkpcbdepsyjhxi
gqykunogzarelkpcbdewsmjhxi
tqyvuwogzarflkpcbdrwsmjmxi
tqdqunogzarflkpcbdewsvjhxi
tqmvunbgzarflkpcbdewsvjhxi
tqyvunogzarflkpcbheesmjhdi
tqxvunogzarfxkpcbdewsmhhxi
tqyvunogzarflkpabdjosmjhxi
tqyvugogztrflkpybdewsmjhxi
tqyvundgzarflkxcbdewsmjhmi
tqyvunogzurfgkpcbdeksmjhxi
tqyvunokzarfckpcbdewsojhxi
tqyvufobzarflkpcldewsmjhxi
tqyvunogznrflkncbdeusmjhxi
tqyvuncgzarfxkpcmdewsmjhxi
oqyvunogzarflkpybdewbmjhxi
tqyvjnogqarfmkpcbdewsmjhxi
tqyvunogzacflkpcidewsmjhwi
tqyvunogzarflkpcbqlwxmjhxi
tqyvunogzarflkpnbhewsvjhxi
vqyiunogzarflkpcbdewsmjhqi
tbyvuncgzarfljpcbdewsmjhxi
tqylunogzarflkpcldexsmjhxi
tqyvulogzarflktcbdewsxjhxi
tqyvmnlgzarflkpcbxewsmjhxi
tqyvunogzartlkpcbdewsmihxp
nqyvunogzarflkpcbdewsmnwxi
tqyvunogzarflkpczdewsmjcxj
tzyvunogzariwkpcbdewsmjhxi
tqyiufogzarflipcbdewsmjhxi
oqyvunogzasflkpcbdewsmjhxv
tqyvunoqmarflkpcwdewsmjhxi
tqrvunogzarflkpqbdewnmjhxi
tqyvunogzarmlkocbdewsmjhri
tqyvunogzakflkpcbveasmjhxi
tqyvunorearflkpcbdewsmfhxi
tqynrnogzarflkpcbdewsmjhxp
tiyvueogzaeflkpcbdewsmjhxi
tqyrunogzarflkpccdewbmjhxi
tqtvunogzarflkpcbdewbnjhxi
tqyvfnogzarflhpcbdewsmjqxi
tqyjunoazarflkpcbdewssjhxi
tqyvunxizarflkpcbdewsmjnxi
tqyvunogzarfhkpobdewsmjhai
tqyaunogzanflkpcbxewsmjhxi
tqyvunogzarflkpsbuewsmjmxi
tqyvunogzarzlkwybdewsmjhxi
tqyvunogzarflkpibdawsmhhxi
tqyvunogzarflkycbdewamjwxi
tqyvunourarflkpcbdewsujhxi
tqyvnnogzirflkpcbdewsdjhxi
tayvunogzauflkpcqdewsmjhxi
tqyvunobzfrflkscbdewsmjhxi
tqygvnogzarflkpcbdewsmjnxi
oqyvunogzarflkpcbdewsmjsgi
tqyvunokzarflkpcbdewshjhii
tqyvunobzarflkvcbdewskjhxi
aqyvunogvarflkpcbdlwsmjhxi
tqyvunogzmrrlbpcbdewsmjhxi
tqyvunggzaqolkpcbdewsmjhxi
tqyvunogqarflkpybdewsmjaxi
tqyvunogzxrflkpcxsewsmjhxi
zqyvunogzarflppcbdewsmjhni
tqvvunogzakslkpcbdewsmjhxi
tqyrunogzarzlkpcbdewsmjtxi
tqyvhnogzarfxkpcbdewqmjhxi
tqyvunogzarflkecbdewgdjhxi
tqyvuwogzerfhkpcbdewsmjhxi
tqmvunogzarflkpcbddwsmdhxi
tqyvunogzarflcqcbdewsmihxi
tqyvunogzarvlkpcbdewsmjmxd
tqyvknogzarfllncbdewsmjhxi
tqyvunogzarflbpcbdrwsajhxi
tqyvunogzarfukpcbddwsmjhii
tqyvuvojzahflkpcbdewsmjhxi
tqyvunogzarfchpcbdeqsmjhxi
wqivueogzarflkpcbdewsmjhxi
tqyvunogzwrflkpcbdewstjhxz
tqyvunogzarfloscbdewsmjhxf
tqfbuiogzarflkpcbdewsmjhxi
tqyvfuogzarflkpcxdewsmjhxi
tqyvunogzarflkpcpdewsmgqxi
tqyvunogzdrflkpcbdewsmqhci
tqyvunogzartlkpcbdewsmjpxj
tqyfunogzarfwkpcbdewsmwhxi
tqyvuntgzarflkjcmdewsmjhxi
tqyvunqfzarflkpckdewsmjhxi
nqyvunogznrflkpcbiewsmjhxi
tqymunobzarflkccbdewsmjhxi
tqyvunogzaoflkprbdewzmjhxi
tqyvunogzaiflkpcvdewbmjhxi
tqwvunogzarfzkpcbdewsmjhxx
txhvunogzarflkpcbdewsijhxi
tqyeunogkarflkicbdewsmjhxi
tqylunogzarylkpcbdewsmkhxi
tqyvriogzarflkpcbdewsmohxi
tqyvunogzsrflkpcbdeasijhxi
tqyvunogzarflkpcbfewcmjhxh
tqyvunoozvrflkpcbdewimjhxi
tqyvunogqayflkpcidewsmjhxi
tqyounogzarflkpccdewsmjhxg
tqgvunogsarflkpcbdewqmjhxi
tqevunogzarflkpcbmewsmjhpi
tqivunogzarflkgcbdewqmjhxi
tqyuunogzlrflkgcbdewsmjhxi
xqyvbnogznrflkpcbdewsmjhxi
tqyvunogzarfjkpebdewsmnhxi
tqyvvnogzarfskpcxdewsmjhxi
thovunogzarflkpcbdewsmjhvi
tqyvunugzarflkpcpdewsmjrxi
tcyvvnogzarflkdcbdewsmjhxi
tqdfunogzarflkpbbdewsmjhxi
tqyvunogzarflkpcbdnwsejzxi
tqyvunivkarflkpcbdewsmjhxi
tqyvunogzawflopcbdewsmjhmi
tqyvunogzarflkpcbdkwsdjqxi
tqyvunodzarflkpcbdewlmjhei
oqyvunoozarflkpcbdemsmjhxi
tiyvunogzarjlkpcbdewfmjhxi
tqrvunogearflkpcbdewsojhxi
tqyvunkgzarflkpcbdcwtmjhxi
tqmvunogzarflkpcbdpwsmjtxi
tqyvunogzarflkpcydeptmjhxi
tqyvunkglarflkpcbdfwsmjhxi
tqyaunogzarflkpcbzeesmjhxi
tqyvunogzarrlkpcbdkwsmjhui
tpyvunogzarflkqcbdewsmjhxd
tvyvunogzarfkkpcbdewsajhxi
gqynunogzarflepcbdewsmjhxi
zqvvunogzarflkpcbdexsmjhxi
tqyyunogzawflkpcbdewsmjhxw
tfyvunogzarflkpcbdewomjrxi
tqyeunogzvrflrpcbdewsmjhxi
nqyvunogzarftkpabdewsmjhxi
tzyvunogzariwkpcbdewmmjhxi
tiyvunogzarflkpcbbewsmjhxa
tqydujoyzarflkpcbdewsmjhxi
tqyvunpgzarflkpcbdeysmjhwi
tqyvunogvarllkpcbdewsmshxi
tqyvunogzbrvlkpcbdewsmjhxp
tcyvueogzarflkacbdewsmjhxi
tqyvunogzrhflkpcbdetsmjhxi
tqavunogzrdflkpcbdewsmjhxi
tqyvunogzjrflkpcbdewstjhki
tqyqunolzarflkpcbdewbmjhxi
tqyvunogzarflkqczdgwsmjhxi
tqyvunogzqrfrkpcbrewsmjhxi
tqyvcnogzhrflkacbdewsmjhxi
tqyvunogzarflkpcbdewsmdhtk
tqyvunoggarftkpcbjewsmjhxi
tvyvunogzarfkkpcpdewsmjhxi
tqyvunogzawflkpcndedsmjhxi
tqyvunogzrrflkpcbdemsmmhxi
tqyvunogzarclkpcbrpwsmjhxi
tryvunogztrflkpcbbewsmjhxi
cqyvundgzarflkpcbdewgmjhxi
tqyvunogzarflktcbkewsmjqxi
tqyvjuogzarflkpcadewsmjhxi
tqyvunogzyrflkpcbbxwsmjhxi
ttyvunogzarflkpcbdewsnmhxi
tqyvunogzarflkpcbdeqsmlhki
fqyvugogzarflapcbdewsmjhxi
tqyvunogzartlkppbdewszjhxi
tqyvunfgztrflzpcbdewsmjhxi
tqyvunogmazflkpcbdewsmjhki
tqyvunzdzarflkpcbdewsmjhvi
tqyvunogzarflkpqbzewsujhxi
tqyvunogzarzlkpcbeewymjhxi
tqyounogzarflkpcbdewsnwhxi
tqysunogsaiflkpcbdewsmjhxi
tqdvunogdarflkpcboewsmjhxi
teyvunogzarflkscbdfwsmjhxi
tqyvunoazarflkpcbdvwsmjhpi
tqyvunooearflkpcbdewrmjhxi
tqyvunoszarflnrcbdewsmjhxi
tqyvunogzalflkpcblewsjjhxi
uqlvunkgzarflkpcbdewsmjhxi
tqyvunobzarflkpcidewsmjhvi
tnyvunogzarflkpcbdnwamjhxi
tqyoudogzarflkpcbdgwsmjhxi
tqyvunoggarflkpcbmewsmwhxi
tqykunogzazflkpcbddwsmjhxi
tfyvunogzarflkpcbdewsmjhgo
tqyvunogztrflkpcbdewoojhxi
tqyvunogzarflkpcbdewbmjoni
tmyvunogzalflkpabdewsmjhxi
tqyvunogzarflkpbbvewqmjhxi
tqyvunofzarflkpcwdexsmjhxi
tayvunogzasflkpcbdewsmhhxi
tqyvlnogzarflkpcbdewsmjwxd
tvyvunogzarflkpcbdewhmjrxi
tqyvunogzarplkpcsdewsmihxi
tqyvunogzarplklcbdewsmjtxi
rqbvunogzarhlkpcbdewsmjhxi
tqyvuxogzarftkpcbdewsmthxi
tqtvunogzarfikpczdewsmjhxi
tqyvunwgzarflkpcbdewsmjxxk
tqygunogzarzlkpnbdewsmjhxi
tqyvunogzarjlkpcbdbwswjhxi
tqyvunogzalfnkpcbdewsmjhxf
tqyucuogzarflkpcbdewsmjhxi
tzyvunogzvrflkpcbdewsmjaxi
tjyvunlgzarflkpcbdewgmjhxi
tqyvcnogzarklkpcbdewsmjhfi
tqyvunogzaaflkpsbaewsmjhxi
tsyvunogzarflkpctdewsmbhxi
tqyeunbgzarflkpcbdewrmjhxi
tqyvunogzcrflqpcbdeismjhxi
eqylunogzarflkpcbdewsmjhxy
tqyvundgzarflkpcbdewmmnhxi
tzyvunogzarflkpcndewsmjhxb
tkyvunogzxrflkpcbdewfmjhxi
tqyvunogzarflkxcbdzwsmjfxi
kqavunogzarflkycbdewsmjhxi
tqyvunogzarjlkpcbdxwkmjhxi
tqyvinogzarfqkpcbdewsmjpxi`
