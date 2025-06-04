
# Wissenszetteltransfer
## Projektauftrag
Auftraggeber: Prof. Dr. Stern

Thema: Interoperabilitätsproblem im Zettelstore

Ausgangssituation: Einzelne Zettel können heruntergeladen werden in verschiedenen Formaten, allerdings nicht als Einheit und können nicht automatisch in andere Zettelkästen eingepflegt werden.

Ziel: Empfänger verfügen über einen erweiterten, aber in sich Konsistenten Zettelkasten mit den angefragten Informationen.

Problem: Es besteht keine Möglichkeit, Wissenszettel in einem eingegrenzten Kontext mit Interessierten Personen zu teilen.

Erwartete Ergebnisse: Der Zettelstore wird erfolgreich um eine Möglichkeit erweitert Wissenszettel zu exportieren.

Rahmenbedingungen:

Zettel im kompatiblen Format vorhanden.
Es müssen keine Metadaten exportiert werden.
Es müssen keine angehängten Dateien exportiert werden.
Dateinamen müssen Unique sein (Kein Merging).  
Abschlusspräsentation: 18.06.2025

Meilensteine:

Erste Datei wird exportiert  
Themenbereiche können markiert werden  
Integration im UI funktioniert  
(Zip-)Komprimierung funktioniert  

Personal:

Eric Nijakowski  Lamia Oktay  Mary Williams  Mia Braun  Dariana Barkov  Stefanie Haag  Melih Akbulut

### Kurzbeschreibung
Wir, als Gruppe 03, Wissenszetteltransfer werden dieses Semester dem Nutzer die Möglichkeit bieten, Wissenszettel in einem eingegrenzten Kontext mit Interessierten Personen zu teilen um den Zettelkasten mit einer Transfermöglichkeit zu erweitern.

## Anleitung

### Vorraussetzung:

go 1.24.3 ist installiert  
Zettelstore 0.20 oder 0.21 vorhanden und unter http://localhost:23123/ erreichbar und läuft im Hintergrund  
Browser entweder Firefox (>=139) oder Chrome (>=137)  
Das Repository wurde bereits heruntergeladen und entpackt oder per `git clone` geklont

### Ausführungsschritte:

ZIP 'Wissenszetteltransfer' Runterladen  
mit 'cd src' den Ordner auswählen  
'go build' eingeben  
'./wissenszetteltransfer.exe' ausführen unter Windows oder Mac: './wissenszetteltransfer'  
'http://localhost:8080/' eingeben im Browser



