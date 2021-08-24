module example.com/main

go 1.17

replace example.com/forexample.com => ../forexample

replace example.com/forexample => ../for

require example.com/forexample v0.0.0-00010101000000-000000000000 // indirect
