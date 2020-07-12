local fp = require('filepath')
local p = require('path')
local base64 = require('base64')
local template = require('template')

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
