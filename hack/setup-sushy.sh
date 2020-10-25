#!/bin/bash

# Make sure there is no VM from previous runs
if virsh list --all --name | grep -qw "vbmc-node"; then
    echo "Cleaning existing vbmc-node..."
    virsh destroy vbmc-node
    virsh undefine vbmc-node
fi

# Create temporary VM definition file
virt-install                \
    --name vbmc-node        \
    --ram 1024              \
    --disk size=1           \
    --vcpus 2               \
    --os-type linux         \
    --os-variant fedora28   \
    --graphics vnc          \
    --print-xml > domain.xml

# Create domain and clean its definition file
virsh define domain.xml
rm -f domain.xml

# Start BMC emulator
sushy-emulator --port 8001 --libvirt-uri "qemu:///system"
