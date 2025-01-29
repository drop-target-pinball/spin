local spin = {}

spin.out = {}

function spin.info(message)
    table.insert(spin.out, { note = {
        kind = 'info',
        message = message,
    }})
end

function spin.play_sound(name, opts)
    table.insert(spin.out, { play_sound = {
        name = name
    }})
end

return spin
