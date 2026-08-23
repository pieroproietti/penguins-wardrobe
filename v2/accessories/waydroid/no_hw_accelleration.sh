#!/bin/sh

sed -i 's/ro.hardware.gralloc=gbm/ro.hardware.gralloc=default/g' /var/lib/waydroid/waydroid_base.prop
sed -i 's/ro.hardware.egl=mesa/ro.hardware.egl=swiftshader/g' /var/lib/waydroid/waydroid_base.prop
