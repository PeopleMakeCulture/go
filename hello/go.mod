module example.com/hello

go 1.20

replace example.com/greetings => ../greetings

replace greetings => ../greetings

require greetings v0.0.0-00010101000000-000000000000 // indirect