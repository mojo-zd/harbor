## Introduction

This guide provides instructions for developers to build and run Harbor from source code.

## Step 1: Prepare for a build environment for Harbor

Harbor is deployed as several Docker containers and most of the code is written in Go language. The build environment requires Python, Docker, Docker Compose and golang development environment. Please install the below prerequisites:


Software              | Required Version
----------------------|--------------------------
docker                | 1.12.0 +
docker-compose        | 1.11.0 +
python                | 2.7 +
git                   | 1.9.1 +
make                  | 3.81 +
golang*               | 1.7.3 +
*optional, required only if you use your own Golang environment.


## Step 2: Getting the source code

   ```sh
      $ git clone https://github.com/vmware/harbor
   ```

## Step 3: Building and installing Harbor

### Configuration

Edit the file **make/harbor.cfg** and make necessary configuration changes such as hostname, admin password and mail server. Refer to **[Installation and Configuration Guide](installation_guide.md#configuring-harbor)** for more info.

   ```sh
      $ cd harbor
      $ vi make/harbor.cfg
   ```

### Compiling and Running

You can compile the code by one of the three approaches:

#### I. Build with offical Golang image

* Get offcial Golang image from docker hub:

   ```sh
      $ docker pull golang:1.9.2
   ```

*  Build, install and bring up Harbor without Notary:

   ```sh
      $ make install GOBUILDIMAGE=golang:1.9.2 COMPILETAG=compile_golangimage CLARITYIMAGE=vmware/harbor-clarity-ui-builder:1.4.0
   ```

*  Build, install and bring up Harbor with Notary:

   ```sh
      $ make install GOBUILDIMAGE=golang:1.9.2 COMPILETAG=compile_golangimage CLARITYIMAGE=vmware/harbor-clarity-ui-builder:1.4.0 NOTARYFLAG=true
   ```

*  Build, install and bring up Harbor with Clair:

   ```sh
      $ make install GOBUILDIMAGE=golang:1.9.2 COMPILETAG=compile_golangimage CLARITYIMAGE=vmware/harbor-clarity-ui-builder:1.4.0 CLAIRFLAG=true
   ```

#### II. Compile code with your own Golang environment, then build Harbor

* Move source code to $GOPATH

   ```sh
      $ mkdir $GOPATH/src/github.com/vmware/
      $ cd ..
      $ mv harbor $GOPATH/src/github.com/vmware/.
   ```

*  Build, install and run Harbor without Notary and Clair:

   ```sh
      $ cd $GOPATH/src/github.com/vmware/harbor
      $ make install
   ```

*  Build, install and run Harbor with Notary and Clair:

   ```sh
      $ cd $GOPATH/src/github.com/vmware/harbor
      $ make install -e NOTARYFLAG=true CLAIRFLAG=true
   ```   
 
### Verify your installation

If everything worked properly, you can get the below message:

   ```sh
      ...
      Start complete. You can visit harbor now.
   ```

Refer to [Installation and Configuration Guide](installation_guide.md#managing-harbors-lifecycle) for more information about managing your Harbor instance.   

## Appendix
* Using the Makefile

The `Makefile` contains these configurable parameters:

Variable           | Description
-------------------|-------------
BASEIMAGE          | Container base image, default: photon
CLARITYIMAGE       | Clarity UI builder image, default: harbor-clarity-ui-builder
DEVFLAG            | Build model flag, default: dev
COMPILETAG         | Compile model flag, default: compile_normal (local golang build)
NOTARYFLAG         | Notary mode flag, default: false
CLAIRFLAG          | Clair mode flag, default: false
HTTPPROXY          | NPM http proxy for Clarity UI builder
REGISTRYSERVER     | Remote registry server IP address
REGISTRYUSER       | Remote registry server user name
REGISTRYPASSWORD   | Remote registry server user password
REGISTRYPROJECTNAME| Project name on remote registry server
VERSIONTAG         | Harbor images tag, default: dev
PKGVERSIONTAG      | Harbor online and offline version tag, default:dev

* Predefined targets:

Target              | Description
--------------------|-------------
all                 | prepare env, compile binaries, build images and install images
prepare             | prepare env
compile             | compile ui and jobservice code
compile_ui          | compile ui binary
compile_jobservice  | compile jobservice binary
compile_clarity     | compile Clarity binary
build               | build Harbor docker images (default: using build_photon)
build_photon        | build Harbor docker images from Photon OS base image
install             | compile binaries, build images, prepare specific version of compose file and startup Harbor instance
start               | startup Harbor instance (set NOTARYFLAG=true when with Notary)
down                | shutdown Harbor instance (set NOTARYFLAG=true when with Notary)
package_online      | prepare online install package
package_offline     | prepare offline install package
pushimage           | push Harbor images to specific registry server
clean all           | remove binary, Harbor images, specific version docker-compose file, specific version tag and online/offline install package
cleanbinary         | remove ui and jobservice binary
cleanimage          | remove Harbor images
cleandockercomposefile  | remove specific version docker-compose
cleanversiontag     | remove specific version tag
cleanpackage        | remove online/offline install package

#### EXAMPLE:

#### Push Harbor images to specific registry server

   ```sh
      $ make pushimage -e DEVFLAG=false REGISTRYSERVER=[$SERVERADDRESS] REGISTRYUSER=[$USERNAME] REGISTRYPASSWORD=[$PASSWORD] REGISTRYPROJECTNAME=[$PROJECTNAME]

   ```

   **Note**: need add "/" on end of REGISTRYSERVER. If REGISTRYSERVER is not set, images will be pushed directly to Docker Hub.


   ```sh
      $ make pushimage -e DEVFLAG=false REGISTRYUSER=[$USERNAME] REGISTRYPASSWORD=[$PASSWORD] REGISTRYPROJECTNAME=[$PROJECTNAME]

   ```

#### Clean up binaries and images of a specific version

   ```sh
      $ make clean -e VERSIONTAG=[TAG]

   ```
   **Note**: If new code had been added to Github, the git commit TAG will change. Better use this command to clean up images and files of previous TAG.

#### By default, the make process create a development build. To create a release build of Harbor, set the below flag to false.

   ```sh
      $ make XXXX -e DEVFLAG=false

   ```
