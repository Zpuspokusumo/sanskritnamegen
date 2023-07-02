package services

import (
	"fmt"
	"math/rand"
)

//make func here for

func sandhify(a string, b string) string {
	isA := false
	isI := false
	isU := false
	//isR := false
	if a == "" {
		return b
	}

	switch string(a[len(a)-1]) { //last element of A
	case "a":
		isA = true
	case "i":
		isI = true
	case "u":
		isU = true
		/* case "ṛ":
		isR = true*/
	}

	switch []rune(b)[0] {
	case 97, 257: //A
		if isA { //Long A
			a = a[:len(a)-1]    //truncate string to -1 of length
			bfirst := []rune(b) //replace first element of string
			bfirst[0] = rune(257)
			b = string(bfirst)
		} else if isI { //YA
			afirst := []rune(a)
			afirst[len(a)-1] = rune(121) // last element from i to y
			a = string(afirst)
		} else if isU { //VA
			afirst := []rune(a)
			afirst[len(a)-1] = rune(118) // last element from u to v
			a = string(afirst)
		}
	case 105, 299: //I
		if isA { //E
			a = a[:len(a)-1]
			bfirst := []rune(b)
			//to e
			bfirst[0] = rune(101)
			b = string(bfirst)
		} else if isI { // II
			a = a[:len(a)-1] //truncate string to -1 of length
			bfirst := []rune(b)
			bfirst[0] = rune(299) // lengthen i
			b = string(bfirst)
		} else if isU { //VI
			afirst := []rune(a)
			afirst[len(a)-1] = rune(118) // last element from u to v
			a = string(afirst)
		}
	case 117, 363: //U
		if isA { //O
			a = a[:len(a)-1]
			bfirst := []rune(b)
			bfirst[0] = rune(101)
			b = string(bfirst)
		} else if isI { //YU
			afirst := []rune(a)
			afirst[len(a)-1] = rune(121) // last element from i to y
			a = string(afirst)
		} else if isU { //UU
			a = a[:len(a)-1] //truncate string to -1 of length
			bfirst := []rune(b)
			bfirst[0] = rune(363) // lengthen u
			b = string(bfirst)
		}
	/*case 7771, 7773:
	if isA { //AR
		bfirst := []rune(b)
		bfirst[0] = rune(82) // change vocalic r to consonant
		b = string(bfirst)
	} else if isI { //YR
		afirst := []rune(a)
		afirst[len(a)-1] = rune(121) // last element from i to y
		a = string(afirst)
	} else if isU { //VR
		afirst := []rune(a)
		afirst[len(a)-1] = rune(118) // last element from u to v
		a = string(afirst)
	}*/

	//case for a-e ai, a-o au,
	default: //do nothing, maybe add consonant to consonant sandhi fix next
	}

	c := fmt.Sprintf("%s%s", a, b)
	return c
}

func Sktnamegen(handom rand.Source) string {
	prefix := []string{"prajna", "maha", "su", "dus"} //a A
	suffix := []string{"in", "ita"}
	/*coolnamesback := []string{"agni", "vana", "ratri", "rupa", "murti", "smrti", "isha"}
	coolnamesback := []string{"mukham", "pani", "manah", "rupa", "murti", "smrti", "isha", "graha", "dvaja"}
	SuObjA := []string{"prajna", "maha", "su", "dus"}
	DusObjA := []string{"prajna", "maha", "su", "dus"} */
	FaceA := []string{"simha", "nara", "agni", "mana", "aga", "megha", "rudhira", "yuddha"}
	// len 10
	BottomA := []string{"jīt", "rūpa", "mukha", "pāṇi", "muṣṭi", "pati", "akṣi", "jata", "cipta", "bhara"}
	var name string
	//length 10
	Nums := []string{"śūnya", "eka", "dvi", "tri", "catur", "pañca", "ṣaṭ", "sapta", "aṣṭa", "nava", "daśa", "sarva"}
	//length 5
	Verbsp := []string{"manta", "jāta", "pata", "datta", "gata", "iccha", "vara", "jaya", "pada", "sarga", "vetta", "varta"}

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

	init := hand.Intn(3)
	// 1 uses prefix, 2 uses num-verbed, 0 skips
	if init == 1 {
		m := hand.Intn(4)
		name = sandhify(name, prefix[m])
		//fmt.Println(m)
	} else if init == 2 {
		m := hand.Intn(11)

		name = sandhify(name, Nums[m])
		m = hand.Intn(8)
		name = sandhify(name, Verbsp[m])
	}

	f := hand.Intn(len(FaceA))
	fnm := FaceA[f]
	b := hand.Intn(len(BottomA))
	bnm := BottomA[b]
	// if string([]rune(fnm)[len(fnm)-1]) == a
	temp := sandhify(fnm, bnm)
	name = sandhify(name, temp)
	//fmt.Println(f)
	//fmt.Println(b)

	init = hand.Intn(2)
	if init == 1 {
		m := hand.Intn(2)
		//if string([]rune(fnm)[len(fnm)-1]) == a
		//fmt.Printf("length %s %s\n", string([]rune(name)[len(name)-1]), name)
		name = sandhify(name, suffix[m])
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
