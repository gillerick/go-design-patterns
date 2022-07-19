package main

func main(){
	listener1 := DataListener{
		Name: "Listener 1",
	}

	listener2 := DataListener{
		Name: "Listener 2",
	}

	subj := &DataSubject{}
	subj.registerObserver(listener1)
	subj.registerObserver(listener2)

	subj.ChangeItem("Monday!")
	subj.ChangeItem("Wednesday!")
	subj.ChangeItem("Thursday!")
}