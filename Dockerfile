from adoptopenjdk/openjdk11
copy easy-vault .
copy easyvault.sh .
user root
CMD ["/bin/bash","-c","./easyvault.sh"]