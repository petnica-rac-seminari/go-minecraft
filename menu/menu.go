package menu

import (
	"main/blocks"
	"os"
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var IsMenu bool = true
var Seed int = 100
var pozadina rl.Texture2D

var prikaziCredits bool = false
var creditsScrollY float32 = 0
var brzinaSkrola float32 = 1.0

var seedTekst string = "100"
var unosAktivan bool = false

const creditsTekst = `CREDITS

Nebo

Mina: Radila na shaderima koji su na kraju propali 🤷.
Igrala odbojku. Napucala slavka, opcinila celo porodicno stablo.
Bila prekrstena od strane Sudomila Prvog,
preskace mine na putu do skole.
Kad se pitala sta je uradila, gledala je okolo
na 2-3 sekunde i rekla idonno.
Popila monstrozne kolicine kafe, jebale ju biblioteke.
Napravila zid-Raca. Polu Jevrej, polu Musliman. Nedjelja.

Nole: Igrao odbojku. Doziveo tesku telesnu povredu
vece pred demo u krvavom fizickom sukobu,
self-proclaimed main character (blam!).
Izvodio orkestru po hodnicima i kafeteriji.
Glumi heteroseksualnost. Lowk uradio vrlo ilegalnu stvar
na sred predavanja koju ne smem da kazem,
i onda se usrao da dolazi policija.

Teodora: Radila na oblacima, menjala gradijent neba.
Snimala svaku glupost koja se desila u Rac-u
i izvan njega u zadnjih 8 dana,
ne zna da implementira osnovne stvari kad ne rade,
krivi team lidera da nije lepo reko kako to treba da funkcionise.

Teona: Zaspala 30 min pred Demo.
Napravila Stefan i Igora, nalazila teksture,
pomogla Teodori za nebo. Investirala u kineske dronove,
imala debatu o ekonomskom stanju u buducnosti.
Ko dodje na petnicu da uci jebeni latinski, kao wtf.
Juri ruse po cetvrtoj.

Voda

Slavica: Ubio se od rada, idalje radi, razboleo se,
zamalo ubio Milutina, rvao se ispred Petnice.
Ucio saradnika kako da kodira.
Doneo cionizam na teritoriju Petnice,
tvrdi da mandarine postanu kisele ako padnu na pod.
Prekrsten od strane Mine-Ban. Polu Jevrej, polu Jevrej.

Marija: Muvala ribu sa fizike.
“Marija sta si ti radila?” “Vodu brt nznm ni ja”
Radila do fora 5 ujutru svaki dan,
ja stvarno nznm kako al kao gg. Nije vibe-codeovala.

Zoe: Implementirala slivanje vode u prostoru,
Paracin znak, bacala hate na Milutina
jer je mazio Petnicke pse,
jadna zavrsila sa fizicarima u sobi.

Grafika

Fikus: Uradio je on tu nesto, vredjao sve bez razloga,
bukv je on jedini pravijo ijednu gresku kao nznm sta prica.
Nebi ni jedan error bio da njega nije bilo,
bukv bi svaki code iz prve runovao bez greske
tacno onako kako treba. Glumio diktatora
(Ne al kao fr on je uradio cetvrtinu svega)

Mihailo: Sjebao sve, pravio gluposti, Co-founder Git-a,
pola prvog dana nije ni programirao nego samo kao
navodno popravljao neke tamo greske 🙄.

Mihajlo: Best friends sa Gemini, Claude-ov secret lover.
(Ne za vibe coding nego reasoning naravno, nisam Dimitrije)
Danju spava, nocu se muva sa arheoloskim ribama.
Ne bi imao snage za rad na igrici bez Purkele i Ustase
lowk goated

Sudomil: Bio nas izvor nade i samopouzdanja,
dao nam snagu da zavrsimo projekat na vreme,
jeo Milutinov toz, postavljao upitne upitke
o upitnim ljudima ili 100 njih.
Bio nosen umesto da radi na projektu,
ali je ok jer ga volimo.

Navigacija

Dejan: Iznervirao se i uzeo pauzu od 4 sata,
za malo izazvao gradjanski rat u svom timu,
radio sve sem za svoj tim

Dimitrije pro: Atleast i still have my vibes,
potencijalan ratni kriminalac. Zaspao 90 min pred demo,
jede mine na putu do skole. Toliko se naradio
da je zadnji dan odmarao. Toliko koristio AI
da se treba smatrati DOS napadom.

Dunja: Napravila pola code koji je Fikus obrisao.
Ne cuje svoj alarm i jedva se budi.
Uradila 60% code-a za basic movement,
followed by crashout i sladoled

Lea: Isto napravila pola code koji je Fikus obrisao.
Osnovala feministicku stranku unutar
demokratske republike Navigacija

Tea: Isto isto napravila pola code koji je Fikus obrisao.
Co-founder FSDRN, govorila svima da smrde,
ucila sve sem racunarstva.

Fikus: Pogledao sav code iz Navigacije,
izbrisao sav code iz Navigacije i napisao novi,
net lines u minusu.

Reljef

Milca: Bila HR. Glumila majku, psihologa,
motivacionog govornika

Palestina: Svadjao se sa urosem, trosio vreme sva 3 dana,
na kraju smo se slagali i mislili na istu stvar
ali nismo iskomunicirali. Pustao sirene da usere noleta

Vrhovni poglavar Ljubisa-Ban: Glumio tatu, Psilohologa,
Project managera, Team managera, moderator zajednickog
komunikacionog medija. Radio matematiku da ostali nebi morali,
court jester, negativan vibe-coding. Pomagao svakom timu,
ladyboy, davao da cure izrazavaju svoju vizuelnu kreativnost
na njegovoj prelepoj faci. Bavio se optimizacijom,
izmislio trigonometriju, founder Sudo-chat,
Project i team organizer. Pravio main menu,
napravio sistem odgovoran za promenu boje neba kroz vreme,
leti preko mina na putu do skole. Polu Musliman,
polu Musliman. Napravio credite (tekst, ne zna da programira kao Mihajlo)

Marko: Igrao minecraft, lowk uradio 90% reljefa,
prekrsten od strane teodore, napravio novu dimenziju.

Honorable mentions:

Teodor: Ljudski pesticid, vodja armije kmetova,
prirodni fenomen, stitio racevce od strasne majke prirode.
Davao predavanje od 2 do 3 ujutru o tome
kako se jedu razve vrste buba, dvocifreni killcount.
Davao predavanje od sat vremena o tome
u sta treba investirati.

Todor: Ljudski herbicid, jeo lisce, karton,
koru od narandze, mandarinu kao jabuku,
za malo da proba i bube da jede.
Blejao sa rac uzivo umesto sa ribom,
ona raskinula sa njim.

Geografija: ulazija na svakih 5 jebenih minuta
da bi pitali da li neko ima cigarete i upaljac,
buraz nemamo kao mi smo racevci zar stvarno mislis
da imamo cigarete jebote kao usli ste 10000 puta
u zdanjih 3 dana koji vam je vise ajte kuci jebote.
bivani kompletno seenovani kad pitjau oce li ko mafije.`

var linijeTeksta []string

func UcitajMenuSliku() {
	pozadina = rl.LoadTexture("menufinal.png")
	linijeTeksta = strings.Split(creditsTekst, "\n")
}

func UnloadujMenuSliku() {
	rl.UnloadTexture(pozadina)
}

func Crtaj() {
	if prikaziCredits {
		rl.ClearBackground(rl.Black)

		creditsScrollY -= brzinaSkrola

		trenutniY := creditsScrollY
		velicinaFonta := int32(20)
		razmakIzmedjuLinija := int32(30)

		for _, linija := range linijeTeksta {
			sirinaTeksta := rl.MeasureText(linija, velicinaFonta)
			centriranoX := (int32(rl.GetScreenWidth()) - sirinaTeksta) / 2

			if trenutniY > -50 && trenutniY < float32(rl.GetScreenHeight()) {
				rl.DrawText(linija, centriranoX, int32(trenutniY), velicinaFonta, rl.White)
			}
			trenutniY += float32(razmakIzmedjuLinija)
		}

		if trenutniY < 0 {
			prikaziCredits = false
		}

		if rl.IsKeyPressed(rl.KeyEscape) || rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			prikaziCredits = false
		}
		return
	}

	rl.ClearBackground(rl.DarkGray)

	izvor := rl.NewRectangle(0, 0, float32(pozadina.Width), float32(pozadina.Height))
	odrediste := rl.NewRectangle(0, 0, float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight()))
	centar := rl.NewVector2(0, 0)
	rl.DrawTexturePro(pozadina, izvor, odrediste, centar, 0.0, rl.White)

	misPozicija := rl.GetMousePosition()

	playRect := rl.NewRectangle(100, 600, 400, 100)
	playBoja := rl.Gray
	if rl.CheckCollisionPointRec(misPozicija, playRect) {
		playBoja = rl.LightGray
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			IsMenu = false
			rl.DisableCursor()
		}
	}
	rl.DrawRectangleRec(playRect, playBoja)
	rl.DrawText("PLAY", int32(playRect.X)+(int32(playRect.Width)-rl.MeasureText("PLAY", 30))/2, int32(playRect.Y)+35, 30, rl.Black)

	creditsRect := rl.NewRectangle(100, 750, 400, 100)
	creditsBoja := rl.Gray
	if rl.CheckCollisionPointRec(misPozicija, creditsRect) {
		creditsBoja = rl.LightGray
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			prikaziCredits = true
			creditsScrollY = float32(rl.GetScreenHeight())
		}
	}
	rl.DrawRectangleRec(creditsRect, creditsBoja)
	rl.DrawText("CREDITS", int32(creditsRect.X)+(int32(creditsRect.Width)-rl.MeasureText("CREDITS", 30))/2, int32(creditsRect.Y)+35, 30, rl.Black)

	quitRect := rl.NewRectangle(100, 900, 400, 100)
	quitBoja := rl.Gray
	if rl.CheckCollisionPointRec(misPozicija, quitRect) {
		quitBoja = rl.Maroon
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			rl.CloseWindow()
			os.Exit(0)
		}
	}
	rl.DrawRectangleRec(quitRect, quitBoja)
	rl.DrawText("QUIT", int32(quitRect.X)+(int32(quitRect.Width)-rl.MeasureText("QUIT", 30))/2, int32(quitRect.Y)+35, 30, rl.Black)

	seedRect := rl.NewRectangle(580, 600, 150, 100)

	if rl.CheckCollisionPointRec(misPozicija, seedRect) && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		unosAktivan = true
	} else if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		unosAktivan = false
	}

	if unosAktivan {
		rl.DrawRectangleRec(seedRect, rl.Black)
		rl.DrawRectangleLinesEx(seedRect, 3, rl.Blue)
	} else {
		rl.DrawRectangleRec(seedRect, rl.Black)
	}

	if unosAktivan {
		kljuc := rl.GetCharPressed()
		for kljuc > 0 {
			if kljuc >= 48 && kljuc <= 57 && len(seedTekst) < 6 {
				seedTekst += string(kljuc)
			}
			kljuc = rl.GetCharPressed()
		}

		if rl.IsKeyPressed(rl.KeyBackspace) && len(seedTekst) > 0 {
			seedTekst = seedTekst[:len(seedTekst)-1]
		}

		if len(seedTekst) > 0 {
			if s, err := strconv.Atoi(seedTekst); err == nil {
				Seed = s
			}
		} else {
			Seed = 0
		}
	}

	if len(seedTekst) == 0 && unosAktivan {
		rl.DrawText("|", int32(seedRect.X)+10, int32(seedRect.Y)+35, 30, rl.White)
	} else {
		rl.DrawText(seedTekst, int32(seedRect.X)+(int32(seedRect.Width)-rl.MeasureText(seedTekst, 30))/2, int32(seedRect.Y)+35, 30, rl.White)
	}

	rl.DrawText("SEED:", int32(seedRect.X), int32(seedRect.Y)-35, 24, rl.White)

	if rl.IsKeyPressed(rl.KeyEnter) {
		IsMenu = false
		rl.DisableCursor()
	}
}

func CrtajHotbar(aktivniBlok blocks.Block) {
	sirinaEkrana := float32(rl.GetScreenWidth())
	visinaEkrana := float32(rl.GetScreenHeight())

	slotovi := [9]blocks.Block{1, 2, 3, 4, 5, 6, 7, 8, 9}
	imena := [9]string{"Water", "Grass", "Dirt", "Stone", "Snow", "Log", "Leaves", "Niggerrack", "Sand"}
	tasteri := [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

	velicinaSlota := float32(70)
	razmak := float32(10)
	ukupnaSirina := (velicinaSlota * 9) + (razmak * 4)

	startX := (sirinaEkrana - ukupnaSirina) / 2
	startY := visinaEkrana - velicinaSlota - 25

	for i := 0; i < 9; i++ {
		rect := rl.NewRectangle(startX+float32(i)*(velicinaSlota+razmak), startY, velicinaSlota, velicinaSlota)

		bojaPozadine := rl.Fade(rl.DarkGray, 0.7)
		bojaOkvira := rl.Gray
		var debljinaOkvira float32 = 2

		if slotovi[i] == aktivniBlok {
			bojaPozadine = rl.Fade(rl.LightGray, 0.8)
			bojaOkvira = rl.Gold
			debljinaOkvira = 4
		}

		rl.DrawRectangleRec(rect, bojaPozadine)
		rl.DrawRectangleLinesEx(rect, debljinaOkvira, bojaOkvira)

		rl.DrawText(tasteri[i], int32(rect.X)+6, int32(rect.Y)+6, 12, rl.White)

		sirinaTeksta := rl.MeasureText(imena[i], 12)
		visinaTeksta := int32(12)
		rl.DrawText(
			imena[i],
			int32(rect.X)+(int32(velicinaSlota)-sirinaTeksta)/2,
			int32(rect.Y)+(int32(velicinaSlota)-visinaTeksta)/2,
			12,
			rl.Black,
		)
	}
}
