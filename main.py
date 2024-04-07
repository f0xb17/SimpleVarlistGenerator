from lib.libfile import LibFile


def main():
    directory = input("Enter the Directory: ")
    option = input("Enter: 1 - varlist / 2 - stringvarlist: ")
    print("--------------------------------------------------------------")
    if option == "1":
        libfile = LibFile(directory, "varlist")

        libfile.findfiles()

        try:
            for file in libfile.getdirectorycontents():
                libfile.writeoutput(file)
                print("Processed: " + file)
                print("--------------------------------------------------------------")
            print("Total amount of variables processed: " + str(libfile.gettotalval()))
        except IOError:
            print("Error while processing the file.")

    elif option == "2":
        libfile = LibFile(directory, "stringvarlist")

        libfile.findfiles()

        try:
            for file in libfile.getdirectorycontents():
                libfile.writeoutput(file)
                print("Processed: " + file)
                print("--------------------------------------------------------------")
            print("Total amount of variables processed: " + str(libfile.gettotalval()))
        except IOError:
            print("Error while processing the file.")
    else:
        print("You're a pretty funny guy or girl.")


if __name__ == "__main__":
    main()
