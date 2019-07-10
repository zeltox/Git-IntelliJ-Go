# Setting up Git, IntelliJ, GO with SSH access to Github in Windows

Make sure latest version of Git is installed for windows. [Download Git](https://git-scm.com/downloads)

## Git Configuration
Add the following lines at global level or project level .gitconfig file depending on the use case.
```
[user]
	name = github_username
	email = email@gmail.com
```

## SSH Configuration
Generate SSH public and private key pairs.
```
ssh-keygen -t ed25519 -N Password_123 -C ssh_key_comment -f /path/to/ssh_key
```

Add your SSH public key to github.com  
Copy your private key to ~/.ssh directory

Add the following lines to ~/.ssh/config file
```
Host github.com
	HostName github.com
	User git
	IdentityFile ~/.ssh/ssh_private_key
```

## Opening IntelliJ
Double clicking on "Open IntelliJ.cmd" will execute start-ssh-agent, which is inside Git installation directory and then opens IntelliJ.  
So now IntelliJ will inherit the SSH_AUTH_SOCK and SSH_AGENT_PID environment variables set by start-ssh-agent.

## Opening Windows Command Line
Double clicking on "Open CMD.cmd" will execute start-ssh-agent and then opens a Windows Command Line.  
Now Windows Command Line will have the same SSH_AUTH_SOCK and SSH_AGENT_PID environment variables as IntelliJ.  
Make sure that your current directory is the project directory for the Git commands to work in the Windows Command Line.

## Notes
As the IntelliJ Command Line and Windows Command Line has the same SSH_AUTH_SOCK and SSH_AGENT_PID environment variables, any SSH private key added to one of them will be available to the other.


## Setting up custom commands
To use the custom commands present in custom_commands directory, set a new PATH environment variable pointing to custom_commands directory.
```
PATH=C:\Users\user_name\Desktop\custom_commands
```
Now these custom commands can be invoked directly from the IntelliJ Command Line or Windows Command Line.  
```ak``` - Add Key (Add private key to the SSH agent)  
```akt``` - Add Key with Time (Add private key to the SSH agent for a specified time duration)  
```akt 3600s```  
```akt 60m```  
```akt 1h```  
```la``` - Lock Agent (Lock Agent the SSH agent with a password)  
```ua``` - Unlock Agent (Unlock Agent the SSH agent with a password)  
```lk``` - List Key (List all private key added to the SSH agent)  
```rk``` - Remove Key (Remove all private key from the SSH agent)  
```ggs``` - Go Get via SSH (Downloads the remote github repository like "go get" but uses SSH)  
```ggs git@github.com:user_name/repo_name.git```  
```ggs https://github.com/user_name/repo_name.git```  
```ggs github.com/user_name/repo_name```  

