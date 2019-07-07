# Setting up Git and IntelliJ IDEA

Make sure latest version of Git is installed for windows. [Download Git](https://git-scm.com/downloads)

## Git Configuration
Add the following lines at global level or project level .gitconfig file depending on the use case
```
[user]
	name = github_username
	email = email@gmail.com
```

## SSH Configuration

Add your ssh public key is added to github.com
Copy your private key to ~/.ssh directory

~/.ssh/config file, Add the following lines
```
Host github.com
	HostName github.com
	User git
	IdentityFile ~/.ssh/ssh_private_key
```

## How "Start IntelliJ IDEA.cmd" works
Double clicking on "Start IntelliJ IDEA.cmd" will execute start-ssh-agent inside Git installation directory and then opens IntelliJ IDEA. 
So now IntelliJ IDEA will inherit the SSH_AUTH_SOCK and SSH_AGENT_PID environment variable set by start-ssh-agent.

## Setting up custom commands
To use the custom commands in custom_commands directory, set a new PATH environment variable pointing to custom_commands directory.
```
PATH=C:\Users\user_name\Desktop\custom_commands
```

Now the custom commands can be invoked directly from the CMD command line.
ak - Add Key (Add Private Key To The Agent)
lk - List Key (List all Private Key Added To The Agent)
rk - Remove Key (Remove all Private Key from The Agent)

