// Finn det minste naturlige tallet som ender på 6 og som har følgende
// egenskap: - hvis man fjerner det siste tallet og plasserer det først så blir
// tallet fire ganger så stort som det opprinnelige tallet.
//
// Løsning: 153846

def shuffle(n: Int): Int = (n.toString.last + n.toString.init).toInt

def valid(n: Int): Boolean = n.toString.last == '6' && shuffle(n) == 4*n

def find(n: Int): Int = if (valid(n)) n else find(n+1)

find(1)
