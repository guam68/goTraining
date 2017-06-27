package main

// goal: map of slices for PAO system
import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	pao := make(map[int]string)

	pao[0] = "Goku"
	pao[1] = "Gohan"
	pao[2] = "Vegeta"
	pao[3] = "Trunks"
	pao[4] = "Piccolo"
	pao[5] = "Frieza"
	pao[6] = "Cell"
	pao[7] = "Buu"
	pao[8] = "Mr Popo"
	pao[9] = "Shenron"

	pao[10] = "Navi"
	pao[11] = "Link"
	pao[12] = "Zelda"
	pao[13] = "Ganondorf"
	pao[14] = "Happy Mask Salesman"
	pao[15] = "Saria"
	pao[16] = "Darunia"
	pao[17] = "Princess Ruto"
	pao[18] = "Tingle"
	pao[19] = "Majora (Imp)"

	pao[20] = "Luke"
	pao[21] = "Leia"
	pao[22] = "Han Solo"
	pao[23] = "Chewbacca"
	pao[24] = "C-3PO"
	pao[25] = "Obi-Wan"
	pao[26] = "Qui-Gon"
	pao[27] = "Queen Amidala"
	pao[28] = "Jar Jar"
	pao[29] = "Anakin"

	pao[30] = "Frodo"
	pao[31] = "Legolas"
	pao[32] = "Gimli"
	pao[33] = "Boramir"
	pao[34] = "Aragorn"
	pao[35] = "Saruman"
	pao[36] = "Uruk"
	pao[37] = "Nazgul"
	pao[38] = "Smeagol"
	pao[39] = "Gandalf"

	pao[40] = "Banjo"
	pao[41] = "Mario"
	pao[42] = "Peach"
	pao[43] = "Bowser"
	pao[44] = "Samus"
	pao[45] = "Fox"
	pao[46] = "Marth"
	pao[47] = "Falco"
	pao[48] = "Ice Climbers"
	pao[49] = "Captain Falcon"

	pao[50] = "Lugia"
	pao[51] = "Ash"
	pao[52] = "Brock"
	pao[53] = "Misti"
	pao[54] = "Prof. Oak"
	pao[55] = "Charizard"
	pao[56] = "Blastoise"
	pao[57] = "Venusaur"
	pao[58] = "Slowbro"
	pao[59] = "Mewtwo"

	pao[60] = "Mickey"
	pao[61] = "Minnie"
	pao[62] = "Donald"
	pao[63] = "Goofie"
	pao[64] = "Pluto"
	pao[65] = "Aladdin"
	pao[66] = "Mulan"
	pao[67] = "Robin Hood"
	pao[68] = "Scar"
	pao[69] = "Mufasa"

	pao[70] = "Revolver Ocelot"
	pao[71] = "Solid Snake"
	pao[72] = "Old Snake"
	pao[73] = "Liquid Snake"
	pao[74] = "Raiden"
	pao[75] = "Psycho Mantis"
	pao[76] = "Sniper Wolf"
	pao[77] = "Vulcan Raven"
	pao[78] = "Otacon"
	pao[79] = "Big Boss"

	pao[80] = "Darth Plagueis"
	pao[81] = "Darth Sidious"
	pao[82] = "Darth Maul"
	pao[83] = "Darth Vader"
	pao[84] = "Grand Moff Tarkin"
	pao[85] = "Storm Trooper"
	pao[86] = "Tie Pilot"
	pao[87] = "Imperial Guard"
	pao[88] = "Jabba"
	pao[89] = "Boba Fett"

	pao[90] = "Master Chief"
	pao[91] = "Cortana"
	pao[92] = "Sgt Johnson"
	pao[93] = "Brute"
	pao[94] = "Prophet"
	pao[95] = "Grunt"
	pao[96] = "Jackal"
	pao[97] = "Elite"
	pao[98] = "Hunter"
	pao[99] = "Flood"

	for {

		rand.Seed(time.Now().UnixNano())
		num := rand.Intn(101)
		reader := bufio.NewReader(os.Stdin)

		fmt.Print(strconv.Itoa(num) + ":")
		answer, _ := reader.ReadString('\n')

		if strings.TrimRight(answer, "\n") == pao[num] {
			continue
		} else if strings.TrimRight(answer, "\n") == "x" {
			break
		} else {
			fmt.Print("Wrong, ")
			fmt.Println(pao[num])
			continue
		}
	}
}
