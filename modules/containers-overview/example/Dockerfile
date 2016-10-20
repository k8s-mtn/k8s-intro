# Thie first instruction in a dockerfile must always be 
# the FROM instruction. This specifies the base image 
# from which you are building. The "scratch" base image 
# is a completely empty base image
FROM scratch 

# MAINTAINER is not required, but is nice to have for reference
MAINTAINER clint@getweave.com

# ADD adds files/directories into the file system of the container.
# For this container I am adding the executable and the assets directory.
ADD webapp /
ADD assets /assets

# Expose informs Docker that the container listens on a 
# certain port. This instruction does NOT expose the port 
# on the host server. To do that, use -p when running 
# the container image.
EXPOSE 8080

# CMD is the command that should be run by default when 
# this container image is started
CMD ["/webapp"]
