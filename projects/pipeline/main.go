package main

func removeDuplicates(inputStream chan string, outputStream chan string) {
	f := " "
	for v := range inputStream {
		if v != f || f == " " {
			outputStream <- v
		}
		f = v
	}
	close(outputStream)
}

func main() {
	channel1 := make(chan string, 3)
	channel2 := make(chan string, 3)
	str1 := "qw"
	str2 := "qw"
	str3 := "jnkj"
	channel1 <- str1
	channel1 <- str2
	channel1 <- str3
	close(channel1)
	removeDuplicates(channel1, channel2)
	for v := range channel2 {
		fmt.Print(v, "; ")
	}

}
