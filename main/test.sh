# !bin/sh
printf "Default tests below:"
printf "\n"
go run . "hello"
go run . "HeLlO"
go run . "Hello There"
go run . "1Hello 2There"
go run . "{Hello There}"

printf "Audit tests below:"
printf "\n"

go run . "hello"
go run . "HELLO"
go run . "HeLlo HuMaN"
go run . "1Hello 2There"
go run . "Hello\nThere\nDipshit" |cat -e
go run . "{Hello & There #}"
go run . "hello There 1 to 2!"
go run . "MaD3IrA&LiSboN"
go run . "1a\"#FdwHywR&/()="
go run . "{|}~"
go run . "[\]^_ 'a"
go run . "RGB"
go run . ":;<=>?@"
go run . "\!\" #$%&'()*+,-./"
go run . "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
go run . "abcdefghijklmnopqrstuvwxyz"

printf "Audit ascii-art-output tests below:"
printf "\n"

go run . "First\nTest" shadow --output=test00.txt
go run . "hello" standard --output=test01.txt
go run . "123 -> #$%" standard --output=test02.txt
go run . "432 -> #$%&@" shadow --output=test03.txt
go run . "There" shadow --output=test04.txt
go run .   "123 -> \"#$%@" thinkertoy --output=test05.txt
go run . "2 you" thinkertoy --output=test06.txt
go run . "Testing long output!" standard --output=test07.txt




printf "<RANDOM> tests below:"
printf "\n"
go run . "fokcYUO"
go run . "dial star 69"
go run . "C++#"
go run . "aa  1**YUH"

printf "Default tests below:"
printf "\n"
go run . "hello" |cat -e
go run . "HeLlO" |cat -e
go run . "Hello There" |cat -e
go run . "1Hello 2There" |cat -e
go run . "{Hello There}" |cat -e

printf "Audit tests below:"
printf "\n"

go run . "hello" |cat -e
go run . "HELLO" |cat -e
go run . "HeLlo HuMaN" |cat -e
go run . "1Hello 2There" |cat -e
go run . "Hello\nThere\nDipshit" |cat -e
go run . "{Hello & There #}" |cat -e
go run . "hello There 1 to 2!" |cat -e
go run . "MaD3IrA&LiSboN" |cat -e
go run . "1a\"#FdwHywR&/()=" |cat -e
go run . "{|}~" |cat -e
go run . "[\]^_ 'a" |cat -e
go run . "RGB" |cat -e
go run . ":;<=>?@" |cat -e
go run . "\!\" #$%&'()*+,-./" |cat -e
go run . "ABCDEFGHIJKLMNOPQRSTUVWXYZ" |cat -e
go run . "abcdefghijklmnopqrstuvwxyz" |cat -e

printf "<RANDOM> tests below:"
printf "\n"
go run . "fokcYUO" |cat -e
go run . "dial star 69" |cat -e
go run . "C++#" |cat -e
go run . "aa  1**YUH" |cat -e

printf "When the last character is a whitespace:"
printf "\n"
go run . "The last character is a whitespace -> " |cat -e

printf "Error output below:"
printf "\n"
go run . "1" "2"

printf "Whitespace tests below:"
printf "\n"
printf "HelloSLASHn"
printf "\n"
go run . "Hello\n" |cat -e
printf "HelloSLASHnSPACeHello"
printf "\n"
go run . "Hello\n Hello" |cat -e
printf "HelloSPACeSLASHnHello"
printf "\n"
go run . "Hello \nHello" |cat -e
printf "SLASHnHello"
printf "\n"
go run . "\nHello" |cat -e
printf "SLASHnHelloSLASHn"
printf "\n"
go run . "\nHello\n" |cat -e
printf "SLASHnSLASHnSLASHn"
printf "\n"
go run . "\n\n\n" |cat -e
printf "SLASHn"
printf "\n"
go run . "\n" |cat -e
printf "Invalid symbols test ❥Hello❥"
printf "\n"
go run . "Invalid symbols test ❥Hello❥" |cat -e
printf "Emty string test:"
printf "\n"
go run . "" |cat -e
printf "SlashNTest:"
printf "\n"
go run . "\n" |cat -e
printf "HelloSlashNTest:"
printf "\n"
go run . "Hello\n" |cat -e
printf "SlashN4timesTest:"
printf "\n"
go run . "\n\n\n\n" |cat -e
printf "spaceSLASHn"
printf " \n"
go run . " \n" |cat -e
printf "SlashN3times+a"
printf " \n"
go run . "\n\n\na" |cat -e
printf "aSlashN3times"
printf " \n"
go run . "a\n\n\n" |cat -e