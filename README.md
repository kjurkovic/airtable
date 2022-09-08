# RAZVOJ WEB APLIKACIJE ZA UPRAVLJANJE KORISNIČKI DEFINIRANIM RELACIJSKIM PODACIMA

Završni rad

Autor: Kristijan Jurković

Mentor: doc. dr. sc. Nikola Tanković

Sveučilište Jurja Dobrile u Puli, Tehnički fakultet


## Sažetak

Razvoj softverskih rješenja se oslanja na dokumentiranju i digitalizaciji postojećih poslovnih procesa. Formulari korišteni u poslovnim procesima modeliraju se prenošenjem postojećih atributa formulara u relacijske baze podataka. Svaka izmjena rješenja zbog izmjene poslovnih procesa mora osigurati konzistentnost i sigurnost postojećih podataka, te zahtjeva određeno vremensko razdoblje kako bi se novi poslovni procesi implementirali.
Jedan od mogućih rješenja su dinamički modeli koji prebacuju modeliranje procesa na krajnjeg korisnika te se programsko rješenje brine o konzistenciji podataka vezanih za definiciju modela (meta model). Nisu potrebne migracije već jednostavno nova definicija modela koja će opisati novi odnosno unaprijeđeni poslovni proces. Nastavno na definiciju modela, prikazan je postupak kreiranja servisa za meta modeliranje i unos podataka pomoću Helm rješenja za menadžment Kubernetes klastera koji su horizontalno skalabilni.

## Funkcionalnosti

* Autorizirani i neautorizirani korisnici
* Autorizirani korisnici:
  * Kreiranje i grupiranje meta modela po radnim okruženjima 
  * Pregled podataka za određeni meta model
* Neautorizirani korisnici:
  * Unos podataka za određeni meta model

Dokumentacija: [Priložena dokumentacija završnog rada](zavrsni_rad.pdf)