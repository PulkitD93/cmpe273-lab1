package main 
import ("testing"
"time")
func TestSleep(t *testing.T) {
	sleep_time:=1
	start_time:=time.Now().Second()
	Sleep(sleep_time)
	end_time:=time.Now().Second()
	t_d := end_time-start_time
	if t_d != sleep_time{
	t.Error("Expected 1, got", t)
}}