<p align="center">
  <img src="Grizzly_logo.jfif" width="350" title="hover text">
</p>


# What is it 

Grizzly is a beta program made to manage your everyday passwords, simple passwords are only supported at this time such as a website and a username. Grizzly impliments a few security features compared to others such as saving your passwords in randomly generated ID files encrypted using AES. The idea of grizzly is to keep hackers from seeing your passwords and not only brute forcing them but also managing to keep the live and decrypted data which is stored in an HTML file hidden and deleted. When you launch the grizzly password manager it will start a web interface on `127.0.0.1://5501`, once you save a password the password is stored in a text file, the text file is of random characters like so `ed@#^&@#^*EDF&*wsft.txt` and the password inside of the file is encrypted using AES, when you go to the passwords tab to view your saved passwords, an HTML template is auto generated, served, and instantly deleted. 

# Current issues and bugs with the project 

The idea of grizzly was to generate the random files, store the locations, and unlock the files in their remote location, this makes it harder for hackers to collect all passwords once your system is comprimised however this has not yet been implimented and is quite buggy, so files are stored in the grizzly filepath. The documentation part of the web interface DOES NOT WORK, it has not been writen yet so please make sure you do not submit a bug report when its simply just not finished. During large passwords like 32 character passwords the password will overflow the HTML and wont load properly.
