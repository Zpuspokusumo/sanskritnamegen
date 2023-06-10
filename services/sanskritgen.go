package services

import (
	"fmt"
	"math/rand"
)

func Sktnamegen(handom rand.Source) string {
	prefix := []string{"prajna", "maha", "su", "dus"} //a A
	suffix := []string{"in", "ita"}
	/*coolnamesback := []string{"agni", "vana", "ratri", "rupa", "murti", "smrti", "isha"}
	coolnamesback := []string{"mukham", "pani", "manah", "rupa", "murti", "smrti", "isha", "graha", "dvaja"}
	SuObjA := []string{"prajna", "maha", "su", "dus"}
	DusObjA := []string{"prajna", "maha", "su", "dus"} */
	FaceA := []string{"Simha", "Nara", "Agni", "Mana", "Aga", "Megha", "Rudhira", "Yuddha"}
	BottomA := []string{"jit", "rupa", "mukha", "pani", "mushti", "pati", "akshi", "jata"}
	var name string
	//Verbsp := []string{"jata", "pata", "datta", "gata", "iccha"}

	/*scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	input := scanner.Text()

	if err != nil {
		fmt.Println("Error reading input. Try again", err)
		return
	}*/

	// handom := rand.NewSource(time.Now().UnixNano())
	hand := rand.New(handom)

	init := hand.Intn(2)
	if init == 1 {
		m := hand.Intn(4)
		name = fmt.Sprintf("%s%s", name, prefix[m])
		//fmt.Println(m)
	}

	f := hand.Intn(len(FaceA))
	fnm := FaceA[f]
	b := hand.Intn(len(BottomA))
	bnm := BottomA[b]
	name = fmt.Sprintf("%s%s%s", name, fnm, bnm)
	//fmt.Println(f)
	//fmt.Println(b)

	init = hand.Intn(2)
	if init == 1 {
		m := hand.Intn(2)
		//if name[len(name)-1] ==
		//fmt.Printf("length %s %s\n", string([]rune(name)[len(name)-1]), name)
		name = fmt.Sprintf("%s%s", name, suffix[m])
		//fmt.Println(m)
	}
	//fmt.Printf("%s\n", name)
	return name

	// tatpurusha, houseborn (loc), tigerthought (gen), sacrifice-knowledge (gen), devadatta (ins)
	// object-verb participle,
	// maybe determine what declension it is

	// karmadharaya-tp, quick-thinking, silvercup, sarvaguna allgood quality,
	// great sage

	// dvigu is just number maker. of the three age, of the three worlds, etc. caturyuga, tribhuvana, etc
	// sahasraaksha hundredeyed

	// nan tatpurusha, negation with a negating 'a' prefix

	// upapada - manuja, svaraaj, ekaja oneborn, jalada (watergiving/er), etc
	// second part is a root

	// aluk tp
	// first element is declined/ uses case
	// dyaus - divaspati (lord of hte sky), atman, atmane pada (forself word)
	// yudh - yudhi s't'hira in battle firm

	// dvandva
	// pANi pAdam handfeet limbs
	// hasty asvas horses and elephants
	// devaa suraas gods and demons

	//bv
	// bahuvriihi, sabretooth
	// indra shatru - indra killer
	// mayuura roman (peacock plumed) ugra baahu strong shoulderd
	// madhu jihva honey tounged
	// vajrapani, vidyahasta, khara mukha

	// dvigu bahuvrihi
	// one wheeled ekacakra
	// caturangga 4 limbed
	// navadvaara ninedoored

	//avyayiibhaava indeclineables, used as adverbials
	// praty-agni, firewards
	// sacakram simultaneously with the wheel
	// antarjalam inter water (within)
	//

	//var seed int32
	//if seed % 3
	// user input here, cool name or a good name

	// determine what type, how long
	// get random name elements
	// fit random name elements according to sandhi
	// sprintf string

	/* var name string
	name =

		fmt.Printf("%s", name)*/
}
