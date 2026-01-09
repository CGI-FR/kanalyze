while true; do
  tr -dc 'A-Z' </dev/urandom | head -c 2
  echo
  sleep .01
done
