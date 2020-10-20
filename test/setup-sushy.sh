#!/bin/bash

if virsh list --all --name | grep -qw "vbmc-node"; then
    echo "Cleaning existing vbmc-node..."
    virsh destroy vbmc-node
    virsh undefine vbmc-node
fi

virt-install                \
    --name vbmc-node        \
    --ram 1024              \
    --disk size=1           \
    --vcpus 2               \
    --os-type linux         \
    --os-variant fedora28   \
    --graphics vnc          \
    --print-xml > domain.xml

virsh define domain.xml
rm -f domain.xml

# Start BMC emulator
sushy-emulator --port 8001 --libvirt-uri "qemu:///system"
