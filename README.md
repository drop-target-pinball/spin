# spin

The Super Pinball System.

Yes, that is a generic name, but it makes for a great package name.
And it is the name of a [great venue in DC](https://wearespin.com/location/washington-dc/).

I am now an owner of a Judge Dredd pinball machine.

![Judge Dredd](judge-dredd.jpg)

It came with a [color dot-matrix display](https://www.colordmd.com/), LED
lighting, and even a secondary power supply. That means it is time to get back
to doing some pinball development work. I am using the
[P-ROC](https://www.multimorphic.com/store/circuit-boards/p-roc/) to control
the machine via USB. I've already blown one fuse by hooking it up wrong.

## Plan

I already have a working version written in Go in the
[main](https://github.com/drop-target-pinball/spin/tree/main) branch. But,
I do need to learn some Rust. And while we are at it, why not change the
game logic to use Lua? This branch is the result of that work.

## Older Demos

Demos from the earlier version. Newer demos on the way!

[![Super Pinball System: alpha-v8](https://img.youtube.com/vi/8MO_zlPVimo/0.jpg)](https://youtu.be/8MO_zlPVimo "Super Pinball System: alpha-v8")

[![Super Pinball System: alpha-v7](https://img.youtube.com/vi/LSEnGz4i4sg/0.jpg)](https://youtu.be/LSEnGz4i4sg "Super Pinball System: alpha-v7")

[![Super Pinball System: alpha-v6](https://img.youtube.com/vi/3ZsQgoLa-z0/0.jpg)](https://youtu.be/3ZsQgoLa-z0 "Super Pinball System: alpha-v6")

## License

MIT