echo --------------------------------------
 echo -e '\033[34mexamples00 6 turns\033[0m'
echo "-------------->"
echo
go run main.go examples/example00.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo -e '\033[34mexamples01 8 turns\033[0m'
echo "-------------->"
echo
go run main.go examples/example01.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo -e '\033[34mexamples02 11 turn\033[0m'
echo "-------------->"
echo
go run main.go examples/example02.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo -e '\033[34mexamples03 6 turns\033[0m'
echo "-------------->"
echo
go run main.go examples/example03.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo -e '\033[34mexamples04 6 turns\033[0m'
echo "-------------->"
echo
go run main.go examples/example04.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo -e '\033[34mexamples05 8 turns\033[0m'
echo "-------------->"
echo
go run main.go examples/example05.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo -e '\033[34mbadexample00\033[0m'
echo "-------------->"
echo
go run main.go examples/badexample00.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo -e '\033[34mbadexample01\033[0m'
echo "-------------->"
echo
go run main.go examples/badexample01.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo -e '\033[34mexamples06 less than 1.5 min\033[0m'
echo "-------------->"
echo
go run main.go examples/example06.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo -e '\033[34mexamples07 less than 2.5 min\033[0m'
echo "-------------->"
echo
go run main.go examples/example07.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
#echo -e '\033[34mexampleMillionAnts less than 2.5 min\033[0m'
#echo "-------------->"
#echo
#go run main.go examples/exampleMillionAnts.txt
#echo
#read -n1 -r -p "Press any key to continue..." key
#echo
#echo --------------------------------------
echo -e '\033[34monly_ants\033[0m'
echo "-------------->"
echo
go run main.go examples/only_ants.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo -e '\033[34mtest_negativeants \033[0m'
echo "-------------->"
echo
go run main.go examples/test_negativeants.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo -e '\033[34mtest_noedges \033[0m'
echo "-------------->"
echo
go run main.go examples/test_noedges.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo -e '\033[34mtest_noend \033[0m'
echo "-------------->"
echo
go run main.go examples/test_noend.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo -e '\033[34mtest_nostart \033[0m'
echo "-------------->"
echo
go run main.go examples/test_nostart.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo -e '\033[34mtest_norooms \033[0m'
echo "-------------->"
echo
go run main.go examples/test_norooms.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo -e '\033[34mtest_toomanyants\033[0m'
echo "-------------->"
echo
go run main.go examples/test_toomanyants.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo -e '\033[34mtest_toomanystart\033[0m'
echo "-------------->"
echo
go run main.go examples/test_toomanystart.txt
echo
read -n1 -r -p "Press any key to continue..." key
echo
echo --------------------------------------
echo
echo "Thank you :)"
exit 0



