package main


import "testing"
/*
basic version
func TestHello (t *testing.T ) {
	got := Hello ("Chris" )
	want := "Hello, Chris"
	if got != want {
		t.Errorf ("got %q want %q ", got, want)
	}

}
*/
/* versione using subtest, more scenario in one function
func TestHello(t *testing.T) {
    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Chris")
        want := "Hello, Chris"

        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    })

    t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
        got := Hello("")
        want := "Hello, World"

        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    })

}*/
/*
Version using helper method
*/
func TestHello(t *testing.T) {

    //assert method
    assertCorrectMessage := func(t *testing.T, got, want string) {
        t.Helper()
        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    }

    //tests
    t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Chris","")
        want := "Hello, Chris"
        assertCorrectMessage(t, got, want)
    })

    t.Run("empty string defaults to 'World'", func(t *testing.T) {
        got := Hello("","")
        want := "Hello, World"
        assertCorrectMessage(t, got, want)
    })

    t.Run("in Spanish", func(t *testing.T) {
        got := Hello("Elodie", "Spanish")
        want := "Hola, Elodie"
        assertCorrectMessage(t, got, want)
    })
    
    t.Run("in French", func(t *testing.T) {
        got := Hello("Elodie", "French")
        want := "Bonjour, Elodie"
        assertCorrectMessage(t, got, want)
    })

}
