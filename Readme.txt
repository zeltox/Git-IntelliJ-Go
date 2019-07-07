# Settingup Git and IntelliJ IDEA

Make sure latest version of git is installed for windows [Download Git](https://git-scm.com/downloads)


.gitconfig file, Add the following lines at global level or prject level depending on the use case
[user]
	name = github_username
	email = email@gmail.com


Add your ssh public key is added to github.com
Copy your private key to ~/.ssh directory

~/.ssh/config file, Add the following lines
Host github.com
	HostName github.com
	User git
	IdentityFile ~/.ssh/ssh_private_key


Double clicking on "Start IntelliJ IDEA.cmd" will execute start-ssh-agent inside Git installation directory and then opens IntelliJ IDEA. 
Then IntelliJ IDEA will inherit the SSH_AUTH_SOCK and SSH_AGENT_PID environment variable set by start-ssh-agent.


To use the custom commands in custom_commands directory, set a new PATH environment variable pointing to custom_commands directory.
Now the custom commands can be invoked directly from the CMD command line.
ak - Add Key (Add Private Key To The Agent)
lk - List Key (List all Private Key Added To The Agent)
rk - Remove Key (Remove all Private Key from The Agent)

