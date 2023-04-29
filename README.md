## Project name: **Lem-in**

The project contains a program written in GoLang, which builds a virtual ant farm according to the provided coordinates in text format, and finds the most optimal paths for ants to get from the "Start" room to the "End" room.

To use the program, you need to have a text file, which contains the number of ants, the list of rooms with their coordinates, and the list of tunnels between the rooms. Please launch the program by using `go run . source.txt` command in your terminal, where "source.txt" should be replaced with the name of your text file.

The pack also contains the file to test the program. To launch it, please just write `bash test.sh` in your terminal if you use Windows, or `./test.sh` for Linux. This will launch a script which will use some pre-installed lists of ant farm rooms to demonstrate how the program works.

Audit page can be found [there](https://github.com/01-edu/public/tree/master/subjects/lem-in/audit)

### **How to read the results?**

The program provides you results in the following format:

```
number_of_ants
the_rooms
the_links

Lx-y
Lx-y Lx-y
Lx-y Lx-y Lx-y
Lx-y Lx-y
Lx-y
```

"**number_of_ants**" is the total number of ants in the farm, which will go from the room "Start" to the room "End". At the beginning, all ants are located in the room "Start", and they can make only one move from room to room in one step.

"**the_rooms**" is the list of rooms with the X and Y coordinates of each room. For example, in the line "A 1 2", "A" will be the name of the room, "1" is the X coordinate of the room, and "2" is the Y coordinate. The list of rooms reflects all rooms in the farm, including the "Start" room and the "End" room marked with special comments.

"**the_links**" is the list of tunnels between rooms. Ants can move within a farm using tunnels from one room to another. The program finds the most suitable tunnels to let all ants come from the "Start" room to the "End" room as soon as possible.

Then there is a **final representation** of each step of each ant, which is done to get from "Start" to "End". `L1` means "the first ant", `L2` is "the second ant", and there are so many ants as it was reflected as the number_of_ants. Each line shows the next steps of ants. At the beginning, all ants are located in the "Start" room. For example, the name of "Start" room is "A", and from this room there are two tunnels - "A-B" and "A-C". In this case, two first ants will make their first steps and it will be displayed in the first line as `L1-B L2-C`. The next lines represent the next steps of the same ants, and the ants who came to the previously used rooms. For example, if the room next to "B" is "D" and the room next to "C" is "E", it can be displayes as `L1-D L2-E L3-B L4-C`, because the new ants came to the tunnels from which the old ants have left. Finally, the scheme represents paths from the "Start" room to the "End" room for all the ants in a farm.

"**Total rounds**" is the total number of lines in the representation, e.g. the number of one-time tunnel changes made by all the ants.

"**Finished in...**" is the timer which shows how fast the calculation was performed.

### **Errors**

What happens if there is an error in the source text file, or there is no option to find paths according to the provided information? You will see an error message in your terminal, if:

* there is an incorrect number of ants (no ants, or invalid data format)
* there is no connection between "Start" and "End" rooms
* there are no detected rooms in the source text file
* there is no detected "Start" room or no detected "End" room
* more than one "Start" or "End" rooms
* tunnels contain rooms that are not represented in the list of rooms
* etc...

The pack contains examples of the text source files, which can be found in the folder "examples". Please feel free to use them.

Created by: **Olena Budarahina** (Gitea username: obudarah), **Kristina Volkova** (Gitea username: Mustkass).