## Ask.fmAnsM
### A simple solution to download Ask.fm Answers' photos and videos without leaving the application..
_________________________________________

#### [Requirements]
- termux -> a terminal emulator for android.
- ask.fm -> android mobile app.
- ROOT (optional but will help)

#### [Setting things up]
###### 1.Install termux from the play store
###### 2.Give termux storage permissions and run the following command
>termux-setup-storage

###### after running those commands to make sure the tool exists ..

>apt update && apt upgrade

###### 3.Compile the program (Ask.fmAnsM) for your architecture (arm/arm64)
to get the binary (about 14+ mbs).

###### 4.Place the binary in your phone's downloads directory.  

###### 5.Go to termux and cd to your downloads directory

>cd $HOME/storage/downloads/

###### 6.Now using termux copy the binary to the app bin directory ..

>cp Ask.fmAnsM /data/data/com.termux/files/usr/bin/

###### now you can run the following command from termux to download the media but still a lot of work !

>Ask.fmAnsM <Answer's link>

#### [Downloading from the app directly]

<img src="https://github.com/ahmdaeyz/Ask.fmAnsM/blob/master/media/1.jpg" width="200" height="200">

<img src="https://github.com/ahmdaeyz/Ask.fmAnsM/blob/master/media/2.jpg" width="200">

<img src="https://github.com/ahmdaeyz/Ask.fmAnsM/blob/master/media/3.jpg" width="200" height="200">

##### To achieve that you need to make a file called termux-url-opener with the following and place it in $HOME/bin

```
#!/bin/bash
 url=$1
 Ask.fmAnsM $url
 mv -f *.png *.jpg *.gif /data/data/com.termux/files/home/storage/pictures/
```

