# Day 1

## Linux Repository Servers
<pre>
- Repository Servers are FTP Servers that maintained multiple opensource installer packages that works on many different Linux versions
- For instance, For Ubuntu Linux distribution, the repository server is maintained by Ubuntu organization
  - for each version of ubuntu os, there are dedicated pages/urls maintained in Repository Servers to provide stable version of software installers
</pre>  

## Linux Package Managers
<pre>
- a utility that comes with every Linux distribution
- used to downad/install/uninstall/upgrade/downgrade softwares
- package manager downloads the softwares from a url maintained in repository configuration files
- the package manager is aware from where to download the software
- the package manager will never download a software that is incompatible to a particular linux distribution
- each Linux distribution supports different package managers
- For example
  - In Debian Linux family, i.e ubuntu, there is a package manager called apt or apt-get, snap, etc.,
  - In Red Hat Linux family
    - Fedora, RHEL, Rocky Linux 
    - package manager supported is dnf, yum, rpm, etc.,
</pre>  

## Info - Provisioning Tools Overview
<pre>
- Infrastructure as a code tool (IaaC)
- these tools helps us declartively create Virtual machine in a on-prem, public or private cloud or in a hybrid cloud
- these tools helps us provision storage in private, public or hybrid cloud environments
- provision switches/routers
- provisioning networks
- examples
  - AWS Cloudformation
  - Terraform
  - Docker
  - Vagrant
</pre>  

## Info - Configuration Management Tool Overview
<pre>
- assumption there is already a machine with some OS installed on it, on a already provisioned machine, configuration management tools can help automate the below activities
- helps us automative system administrative activities
  - software installation
  - software uninstallation
  - update/upgrading/downgrading softwares
  - managing users
    - creating a windows/linux/unix/mac user
    - editing a windows/linux/unix/mac user rights
    - deleting a windows/linux/unix/mac user
    - updating user details, credentials, etc.,
  - adding a windows/linux/unix server to a domain
  - removing a windows/linux/unix server from a domain
    - installing software patches
- managing switches/routers
- creating a vlan, editing, updating network stuffs
- examples
  - Puppet
  - Chef
  - Salt/Salt Stack
  - Ansible
</pre>

## Info - Puppet Overview
<pre>
- is one of the oldest configuration management tool
- uses its own proprietary declarative language called Puppet language to write automation scripts
- follows client/server architecture
- DSL(Domain Specific Language) - language used to write the automation 
- DSL - Puppet Language
- follows Pull based architecture
- installing Puppet is a complex activity
- learning curve is also steep
- a special Puppet agent must be installed on all servers that must be managed by Puppet
- comes in two flavours
  - opensource
  - enterprise product
</pre>

## Info - Chef Overview
<pre>
- chef is one of the very popular configuration management tools available in the market
- follows client/server and pull based architecture
- DSL used in Chef is Ruby scripting language
- installation and learning is quite difficult
- Chef comes with many tools, hence mastering chef takes very long time
- architecture wise, somewhat similar to Puppet
- a special Chef agent must be installed on all servers that must be managed by Chef
</pre>

## Info - Ansible Overview
<pre>
- is the youngest configuration management tool 
- it is developed in Python by Michael Deehan
- Michael Deehaan is a x employee of Red Hat
- Michael Deehan started a company called Ansible Inc, Ansible core was developed as an open source product
- Ansible follows a very simple architecture
- Installation of Ansible is very easy
- Easy to learn and master Ansible
- Ansible is agent-less
- Ansible comes in 3 flavours
  1. Ansible Core - opensource product, supports only command line interface
  2. AWX - open source product developed on top of Ansible Core, supports Web Interface
  3. Red Hat Ansible Automation Platform ( earlier called as Ansible Tower )
     - developed on top of the open source AWX
     - supports web interface
     - also can expect world-wide support
     - Red Hat is an IBM company
</pre>

## Info - Ansible High Level Architecture 
![Ansible](AnsibleHighLevelArchitecture.png)
