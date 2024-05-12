# Ascii-art Generator
-Ascii-art generator is a golang program tha converts user input into ascii-art.
The program mainly takes user input from the terminal and then prints the output in ascii format in the same terminal.
## Description 
<p>the program is coded with golang only and uses the standard golang packages<p>
<p>The main file has the main.go file and the ascii directory with three text files which are the banner files.<p>
<P>the banner text files which  are shadow.txt, standard.txt and thinkertoy.txt are the file containing the prepared ascii art values matching all printable values each ordered as they are in the ascii table, they are all of different values and when used they produce different ascii art representation.<p>
<p>the ascii directory contains two files which are the asciiart.go and filereader.go which hold fucntions which are used  to print and to read the file<p>

## how it works
<P>the user must initate a terminal in the ascii-art directory.<p>
<p>the command to run is " go run .{this is main program, you can use main.go instead of ' . '}     -t|-s   {optional but for different bannertext}  {userinput in quotes}<p>
<p>one has an option case of using "flags" with two optional flags: -t for the thintertoy style and -s for the shadow style. In case the flags are not used, then the standard style is displayed. <p>
<P>The user can put the input in double quotes or without double quotes, In the case of without double quotes, the user must escape all bash special characters <p>
<p>To display your text separately, put the character "\n" like so, "Hello\nthere". THe program will display hello and there on separate line fields. <p>
<p>To edit your text to display in tabs, use "\t", vertical tabs use "\v", form feed use "\f", backspace use "\b", carriage ret use "\r" and Delete use "DEL"<p>
<P>All edit characters should be escaped in cases where no quotes is used eg \\t<P>
<p>If the command is run with only one input that does not have quotes, the program will execute it, otherwise it will print an error.<p>

## Credit
<p>The program was written and is being maintained by Fena Onditi, James Muchiri and Jesee Kuya.<p>



