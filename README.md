# yaml2dic
yaml file to dictionnary

## build
    go build

## example
given this yaml file

    # Define a key-value pair
    name: John Doe

    # Define a nested key-value pair
    address:
      street: 123 Main Street
      city: Anytown
      state: Any State
      postal_code: 12345

    # Define a list of values
    hobbies:
      - hiking
      - reading
      - cooking

parse it

    $ cat example.yaml | ./yaml2dic
    .hobbies=[hiking reading cooking]
    .name=John Doe
    .address.street=123 Main Street
    .address.city=Anytown
    .address.state=Any State
    .address.postal_code=12345

# dowload it from release

    wget https://github.com/thibault-ketterer/yaml2dic/releases/download/1.0.0/yaml2dic.tar.gz
    tar xf yaml2dic.tar.gz
    mv yaml2dic ~/bin/.

