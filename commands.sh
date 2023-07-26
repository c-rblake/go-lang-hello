# GIT create a repository from the command line. needs the CLI
# After you've created the repository on GitHub, you can push your local Git repository to GitHub using the git push command. Here's how
#git remote add origin https:#github.com/c-rblake/go-lang-hello
#git branch -M main
# git push -u origin main

# Create a local branch with INIT before creating the remote repository. Then has to be solved
# git log --graph --oneline --all  # to see the graph
# To solve this, you can use the --allow-unrelated-histories option with the git merge command:

# go run main.go hello.go -- will compile and run both files
# go run main.go hello.go --help -- will show help
# go build -o myprogram.exe compiles all the go files to one Executable. NEAT!
# ./myprogram

# go get -u github.com/gorilla/mux

# Powershell build alias. 
# # Set-Alias -Name build -Value ~\build.ps1   (first add go build -o myprogram.exe) to the build.ps1 file


# Add a GO routine