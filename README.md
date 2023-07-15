# SANSKRIT NAME GENERATOR GUI APP
#### Video Demo:  <https://youtu.be/e9UQtBBjWSk>
### EDIT: I have moved the API code to server/servermain.go, and this main.go file will contain the code to the GUI version of this app. My grandma was the antonym of helping in the development of this app, so whatever but shes still my grandma
#### Description: Sanskrit name generator GUI APP by zpuspokusumo
#### I built it in Go using Gin Framework. It generates a random sanskrit names using a random number generator to determine how long and how what prefix and/or suffix, if any, it will have. It will then generate a random number to index one of many slices, which will then be joined using the Sandhify function to be grammatical.
#### My name is Zidane, i am from a non IT background. I graduated from Universitas Airlangga with an English Linguistics degree in June of 2022, after which i took CS50 just because i didnt know what i wanted to do, but has always been curious about programming.I finished week 10 on December of 2022, and it took me until now, around 6 months, to come up with a final project that i would use and be proud of creating. I even got a job between that and now, which i unfortunately could not keep. However, it has been a great learning experience nontheless, and does not deter me from further diving into the tech industry as i believe i have a great understanding of the fundamentals of computer science.
#### For this project, i combined sanskrit, a language i enjoyed learning, with Go, another language i enjoy learning. The main.go file contains the api logic, and the services folder with the sanskritgen.go file contains the random sanskrit name generation algorithm. I still think it can be improved and expanded, but for now, i am satisfied with the selection of word-elements for the names, and the Sandhify function.
#### The Sktnamegen function generates names, and calls the Sandhify() function to join them together according to sandhi. As of now it can join vowels according to sandhi, and i intend to expand it to also be able to join and generate geminated consonants and Vrddhi vowels as well.
#### I hope that this project can similarly inspire you to learn how to code, the Go language, or even learning sanskrit. CS50 was a remarkable and incredible experience, that i felt was beneficial especially to tech-curious demographic who are unable to learn about computer science concepts outside of a bootcamp, even if said bootcamp actually discusses those topics as in depth as they are discussed in CS50X. I have yet see any bootcamp that offers a course in C, or even a memory handling section. CS50 is neither a bootcamp nor is it comparable to one, it is a full on education, now accessible through the internet for everyone to see and learn at no cost, and you can get a free certificate to boot. There is never enough thanks from me and people like me, whose lives have been similarly changed by this lifechanging experience, to David J Malan our instructor, Brian Yu, Doug Lloyd, and the rest of the CS50 staff for being incredible human beings and inspiring countless people. I even have plans to teach my local neighborhood kids how to code for free as well, however, i will do so after i do not have to worry about getting a job.
### how to use (for now):
#### 1. run the Sanskritgen.exe
