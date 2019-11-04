###### Konkurentné matice ######

Strategia klienta na rozdelovanie uloh serverom je zalozena na jednoduchej synchronizacii. V metode mainLoop workerPoolu sa na zaciatku vytvori
buffered channel "counter" s velkostou len(wp.workers) a nasledne pri priradovani kazdej ulohy do tohtou channelu napiseme nejaku hodnotu, cim si zaznacime,
ze ju bude niektory s workerov riesit. Tato hodnota je z channelu odobrata az vtedy, ked bola dana uloha dokoncena, to znamena, ze si tak evidujeme
pocet volnych workerov a ked nie je nikto volny (channel je plny), workerPool stoji na zapise counter <- true, az pokym niekto nedokonci ulohu a 
nespravi _ = <- counter...

Meriame aj celkovy "pracovny" cas klienta a serverov. Po vyrieseni aktualnej ulohy (znasobeni matic) klient informuje servery, ze "brigada" je hotova,
a ti mu nasledne poslu pocet milisekund, kolko za pocas nej "namakali". Na vymenu tejto spravy je zneuzita struktura Result/Task. Klient jednoducho 
posle &Task{-1, Vektor{}, -1, Vektor{}} (-1 na mieste indexov pre vektory), a server vie, ze na takuto spravu ma odpovedat prave poctom milisekund,
takze posle naspat 7Result{-1, -1, Cislo(workTime.Milliseconds())}. Mozno trocha hack, no prislo mi zbytocne riesit vymenu nejakeho ineho typu 
sprav, kedze aj do tohto sa da zaobalit potrebna informacia. 

Ukazka vypisu na strane klienta (N=100):

Z celkoveho casu 3.2767381s client pracoval 7.9993ms
Takze jeho vytazenost bola 0.2136752136752137 percent

Zistujem vytazenie serverov...
worker 8001 pracuje
Server 192.168.100.13:8001 pracoval celkovo 781ms, takze jeho vytazenost bola: 23.84004884004884 percent
worker 8002 pracuje
Server 192.168.100.13:8002 pracoval celkovo 545ms, takze jeho vytazenost bola: 16.636141636141634 percent
worker 8003 pracuje
Server 192.168.100.13:8003 pracoval celkovo 528ms, takze jeho vytazenost bola: 16.117216117216117 percent

Vidime, ze klient sa velmi "nenamakal" a vacsinu casu iba cakal na pracu serverov - to sme cakali. Prave vdaka tomu bolo vhodne zvolit
sychronizaciu cez channel a nie aktivne cakanie v nekonecnom loope, kedze by sme tak zabili vela CPU casu. Dalej si mozno vsimnut, ze sucet
casov jednotlivych serverov sa nerovna celkovemu casu behu - zbytok casu bol zjavne straveny pri samotnom prenose dat, co ale nepocitame
do pracovneho casu klienta ani servera.

No a prikladam aj fotku toho (matice/ukazka.jpg), ako to vyzeralo v akcii.

######### Determinant ##########

O tejto ulohe pisem dost v komentaroch pri kode. Naivna paralelizacia cas vypoctu rapidne zhorsuje, pretoze sa vytvori neumerne
vela go rutin (DeterminantParallel). Po sikovnom obmedzeni (DeterminantLimiter) to uz ale dava slusne vysledky!

########### Crawler ############

Respektoval som interface Fetcher a metoda Fetch preto vracia aj telo nacitanej stranky ako string, avsak v samotnych crawleroch tento vystup
ignorujem a zaujimam sa len o zoznam url adries, kedze mi to prislo zbytocne a vypisovanie celeho body zbytocne spomaluje a skaredi vypis...

Moj RealFetcher nie je ziadne terno, jednoduchy regex, ktory najde linky zacinajuce sa na href (kod z prednasky), a potom si z toho este vysekne
linkovu cast. Obcas sa sekne a dostaneme nejake to 40X chybu, ale to nam v zasade nevadi, ten link proste preskocime... 
Nejake statistiky (depth = 2):

CrawlBasic:
  http://dai.fmph.uniba.sk/courses/PARA
  size: 112
  4m1.5760391s

  http://golang.org/
  size: 116
  1m23.6363422s

CrawlRecursive
  http://dai.fmph.uniba.sk/courses/PARA
  size: 52
  28.4810503s

  http://golang.org/
  size: 85
  55.354704s

CrawlConcurrent
  http://dai.fmph.uniba.sk/courses/PARA
  size: 112
  1.197505s

  http://golang.org/
  size: 103
  2.4180273s

Doimplementoval som vsak este jeden Fetcher, ktory nehlada iba linky zacinajuce href=, ale skratka vsetko, co vyzera ako link.
Jeho vysledky:

CrawlBasic:
  http://dai.fmph.uniba.sk/courses/PARA
  size: 224
  6m46.7572347s

  http://golang.org/
  size: 132
  1m45.566631s

CrawlRecursive
  http://dai.fmph.uniba.sk/courses/PARA
  size: 73
  35.3066808s

  http://golang.org/
  size: 88
  51.4587789s

CrawlConcurrent
  http://dai.fmph.uniba.sk/courses/PARA
  size: 229
  1.4727397s

  http://golang.org/
  size: 132
  1.8358456s

Zahadu preco davaju crawlery navzajom rozne pocty najdenych stranok som nevyriesil... V kode nic podozdrive nevidim, kontroloval som podmienky na 
depth, no pride mi to vsade okej... Jedine co ma napada je to, ze navstevuju jednotlive stranky v roznom poradi. Je mozne, ze niektore im daju
40X error ked daju prilis vela requestov po sebe, pricom iny crawler na toto nemusi narazit, pretoze ide "inymi" vetvami. Ale to je len taky tip...

################################