# Reminder

## Application to make reminders for all you need 

![reminder](./.github/images/reminder.gif)

## Usage

```shell
$ go run main.go -timer=10m -text="Water time!!"
```

Or can you build this project and add on your bin path like that:

**MacOs and Linux users**

```shell
$ go build -o reminder
$ sudo cp reminder /usr/local/bin/reminder
$ reminder -timer=10m -text="Water time!!" &
```

**Windows users**

You are not welcome here :D

## Flags 

The flag `-timer` is about the time that alert will happens, for example, `10m` you alert will happen each 10 minutes.

The flag `-text` is about the text that will show in your alert.

Default values to `-timer` and `-text` is `10m` and `Reminder time` respectively

## Alert sample

When you executed that program you will receive an alert like that
![alert](.github/images/alert.png)
