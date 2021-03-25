# UdemyGo
This is my UdemyGo Project

	Note from course:

		- How to run go code ?
			A:  terminal-> go run main.go
			    go build - just compile (new file)
			    go run - compile + exectude 

		- What does "package main" mean ?
            A:  Package is like a Project/Workspace
                Eatch file depending on the package xxx needs to start with package xxx
                There are 2 types of packages
                Exectuable:
                    when compiles spits out a executable file
                and Reusable:
                    libs / "helper" / reusable logic
                    for future projects
                
                The name of the package determinates what type you are using. Especially the word "main" = executable
        
        - What does "import fmt" mean ?
            A:  Works like "import" works in other languages. golang.org/pkg/ for standard libs

        - What does "func" mean ?
            A:  Works like "functions" works in other languages. 

        - If you first declare a var use :=

        - Array vs Slice
            Silce is "basically" a Array with more features
            Array is fix length - Silce can grow and shrink

        - func (d deck) print(){
            Set a methode to a variable we created !
            Any variable of type deck get access to the message print()
            d is the actually reference to the 
            variable calling the print()

        - Array[0:3]
            Could be written as Array[:3] its 0,1,2 - so 3 is not included.
            Array[3:] is the rest of array till the end

        - type conversion
            []byte("Heyo")