local spin = {}

spin.conf = {}
spin.out = {}

function spin.alert(message)
    table.insert(spin.out, { note = {
        kind = 'alert',
        message = message,
    }})
end

function spin.fault(message)
    table.insert(spin.out, { note = {
        kind = 'fault',
        message = message,
    }})
end

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

function spin.run(name)
    table.insert(spin.out, { run = {
        name = name
    }})
end

package.loaded['spin'] = spin

