#!/bin/bash

# ----------------------------------------------------------------------
# Configuración de LightDM para Quirinux
# Ejecutado por tailor durante el finalize del wardrobe.
# ----------------------------------------------------------------------

# ----------------------------------------------------------------------
# 1. Determinar usuario objetivo
# ----------------------------------------------------------------------

if [ -n "$DOAS_USER" ] && [ "$DOAS_USER" != "root" ]; then
    TARGET_USER="$DOAS_USER"
elif [ -n "$SUDO_USER" ] && [ "$SUDO_USER" != "root" ]; then
    TARGET_USER="$SUDO_USER"
else
    TARGET_USER="$(awk -F: '$3 >= 1000 && $6 ~ /^\/home\// && $7 !~ /(nologin|false)/ { print $1; exit }' /etc/passwd)"
fi

if [ -z "$TARGET_USER" ]; then
    TARGET_USER="charlie"
fi

echo "[config_lightdm] Usuario objetivo: $TARGET_USER"

# ----------------------------------------------------------------------
# 2. Determinar archivo de configuración de LightDM
# ----------------------------------------------------------------------

mkdir -p /etc/lightdm/lightdm.conf.d

LIGHTDM_CONF="/etc/lightdm/lightdm.conf.d/99-quirinux.conf"

echo "[config_lightdm] Archivo de configuración: $LIGHTDM_CONF"

# ----------------------------------------------------------------------
# 3. Configuración de LightDM
# ----------------------------------------------------------------------

cat > "$LIGHTDM_CONF" <<EOF
[Seat:*]
# Quirinux wardrobe
autologin-user=${TARGET_USER}
autologin-user-timeout=0

# Forzar sesión XFCE para login manual y autologin.
# Evita que LightDM use el fallback lightdm-xsession.
user-session=xfce
autologin-session=xfce

allow-user-switching=true
allow-guest=false
greeter-session=lightdm-gtk-greeter
EOF

echo "[config_lightdm] LightDM configurado para sesión XFCE"

# ----------------------------------------------------------------------
# 4. Fondo del greeter
# ----------------------------------------------------------------------

COSTUME="${1:-quirinux}"
BACKGROUND_DIR="/usr/share/backgrounds/${COSTUME}"

if [ -d "$BACKGROUND_DIR" ]; then
    BACKGROUND="$(ls "$BACKGROUND_DIR"/*.jpg 2>/dev/null | head -1)"

    if [ -n "$BACKGROUND" ]; then
        touch /etc/lightdm/lightdm-gtk-greeter.conf

        if grep -q "^background=" /etc/lightdm/lightdm-gtk-greeter.conf 2>/dev/null; then
            sed -i "s|^background=.*|background=${BACKGROUND}|" /etc/lightdm/lightdm-gtk-greeter.conf
        else
            printf "\n# wardrobe\nbackground=%s\n" "$BACKGROUND" >> /etc/lightdm/lightdm-gtk-greeter.conf
        fi

        echo "[config_lightdm] Fondo configurado: $BACKGROUND"
    fi
fi

# ----------------------------------------------------------------------
# 5. Forzar sesión XFCE en /etc/skel/.dmrc
# ----------------------------------------------------------------------

mkdir -p /etc/skel
printf "[Desktop]\nSession=xfce\n" > /etc/skel/.dmrc
chmod 644 /etc/skel/.dmrc

echo "[config_lightdm] /etc/skel/.dmrc configurado"

# ----------------------------------------------------------------------
# 6. Forzar sesión XFCE en usuarios existentes
# ----------------------------------------------------------------------

for h in /home/*; do
    [ -d "$h" ] || continue

    u="$(basename "$h")"

    id "$u" >/dev/null 2>&1 || continue

    uid="$(id -u "$u" 2>/dev/null || echo 0)"
    [ "$uid" -ge 1000 ] 2>/dev/null || continue

    if [ -f "$h/.dmrc" ]; then
        if grep -q "^Session=" "$h/.dmrc"; then
            sed -i "s/^Session=.*/Session=xfce/" "$h/.dmrc"
        else
            printf "\nSession=xfce\n" >> "$h/.dmrc"
        fi

        if ! grep -q "^\[Desktop\]" "$h/.dmrc"; then
            sed -i '1i[Desktop]' "$h/.dmrc"
        fi
    else
        printf "[Desktop]\nSession=xfce\n" > "$h/.dmrc"
    fi

    chown "$u:$u" "$h/.dmrc" 2>/dev/null || true
    chmod 644 "$h/.dmrc" 2>/dev/null || true

    echo "[config_lightdm] $h/.dmrc configurado para $u"
done

# ----------------------------------------------------------------------
# 7. Corregir AccountsService
# ----------------------------------------------------------------------

if [ -d /var/lib/AccountsService/users ]; then
    for f in /var/lib/AccountsService/users/*; do
        [ -f "$f" ] || continue

        if grep -q "^XSession=" "$f"; then
            sed -i "s/^XSession=.*/XSession=xfce/" "$f"
        else
            printf "\nXSession=xfce\n" >> "$f"
        fi

        if grep -q "^Session=" "$f"; then
            sed -i "s/^Session=.*/Session=xfce/" "$f"
        fi

        echo "[config_lightdm] AccountsService corregido: $(basename "$f")"
    done
fi

# ----------------------------------------------------------------------
# 8. Asegurar que LightDM sea el display manager
# ----------------------------------------------------------------------

if [ -x /usr/sbin/lightdm ]; then
    echo "/usr/sbin/lightdm" > /etc/X11/default-display-manager
    echo "[config_lightdm] LightDM establecido como display manager"
fi

# ----------------------------------------------------------------------
# 9. Deshabilitar otros display managers y habilitar LightDM
# ----------------------------------------------------------------------

if command -v systemctl >/dev/null 2>&1; then
    systemctl set-default graphical.target >/dev/null 2>&1 || true
    systemctl enable lightdm.service >/dev/null 2>&1 || true

    for dm in slim gdm3 gdm sddm; do
        systemctl disable "${dm}.service" >/dev/null 2>&1 || true
        systemctl mask "${dm}.service" >/dev/null 2>&1 || true
    done

    echo "[config_lightdm] LightDM habilitado vía systemd"
else
    if [ -x /usr/sbin/update-rc.d ]; then
        for dm in slim gdm3 gdm sddm; do
            if [ -f "/etc/init.d/${dm}" ]; then
                update-rc.d -f "$dm" remove >/dev/null 2>&1 || true
            fi
        done

        update-rc.d lightdm defaults >/dev/null 2>&1 || true
        update-rc.d lightdm enable >/dev/null 2>&1 || true

        echo "[config_lightdm] LightDM habilitado vía sysvinit"
    fi
fi

echo "[config_lightdm] Configuración completada"
exit 0
