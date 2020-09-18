local fp = require('filepath')
local p = require('path')
local base64 = require('base64')
local template = require('template')
local env = require('env')

print(fp.from_slash(p.join('a','b','c')))
print(fp.split_list("a;b;c")[1])

local content = base64.encode_file(fp.join('base64.go'))

local dec, err = base64.decode(content)
if not dec then
    print(err)
else
    local content2 = base64.encode(dec)
    if content ~= content2 then
        print('content do not match')
    end
end


local rendered = template.render([[
    Hello dear {{ .person.name }}, welcome to the wonders of glua!
]], { person= { name= "Bob Lazar"}})
print(rendered)


-- windows doens't like files named aux.......
local f, err = io.open("auxiliar.txt", "wb")
if err then
    print(err)
    os.exit(1)
end
f:write('hello')
f:close()


local execute = require('exec').execute
ok, exitCode, err = execute("go", "version")
if not ok then
    print("ec", exitCode, "err", err)
end

ok, exitCode, err = execute("this-binary-does-not-exist")
print(ok, exitCode, err)


env.set("HI", "WORLD")
if env.get("HI") ~= "WORLD" then
    print("unable to set env value")
else
    print("value is OK")
end
