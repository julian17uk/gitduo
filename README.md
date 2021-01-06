# Gitduo

This is small CLI to help maintain multiple GitHub accounts and repository creation.

## Getting Started

Clone this repo and try building the software in cmd directory with  

$go build -o gitduo  

Then copy the executable into the bin directory with  

$cp gitduo /Usr/local/bin/gitduo  

Any NEW terminal will now have access to the CLI

Run a test with  
$gitduo help

### Prerequisites

This required a unix terminal to operate the CLI  
Golang is required  
Two working github accounts are needed to operate with SSH key pairs, not HTTPS  
A GitHub Personal Access Token is needed for repository creation.


#### Setup github accounts

If you already have a github account on your machine with ssh access you should have a id_rsa and id_rsa.pub files in your ~/.ssh directory.

For the second github user we need to be careful not to overwrite these files.

Move to the ~/.ssh directory  
$cd ~/.ssh

Create new keys  
$ssh-add ~/.ssh/id_rsa_work

Edit config file  
$vi config

Original contents will look something like this:  
Host github.com  
  AddKeysToAgent yes  
  UseKeychain yes  
  HostName github.com  
  User git  
  IdentityFile ~/.ssh/id_rsa  

We need to add a second host like this:  
Host github-work  
  Hostname github.com  
  User git  
  IdentityFile ~/.ssh/id_rsa_work  

Whilst you are still in the ~/.ssh directory we need to create a couple of more files  
$touch gittoken  
$touch gitconf

Add the second users github personal access token (pat) to the gittoken file. To test it is there you can run   
$gitduo pat

We also need to add the configurations to gitconf file in json format.
Please use the gitconf.sample provide as a template and note that the Host remains unchanged for both Main and Work

Further you need to grant access from your main github account repo to the work github account for this and vice versa for this to work properly

It is possible to change the location (and name) of the gittoken and gitconf files. They are speficied as variables in main.go see here

var PersonalAccessTokenDirectory = "/.ssh"  
var PersonalAccessTokenFileName = "gittoken"  
var ConfUserDataDirectory = "/.ssh"  
var ConfUserDataFileName = "gitconf"  

#### Deployment

Try out gitduo with the following commands:

$gitduo help  
$gitduo pat

And inside a git directory

$gitduo which  
$gitduo main  
$gitduo work

Or create a new repository from sratch in an empty directory with the following  
$gitduo set 1  (for a private github repo)  
$gitduo set 0  (for a public github repo)