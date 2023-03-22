print("hello lua")

-- var
local x = 10
local name = "min"
local flag = true
local a = nil 

-- plug
local n = 1
n = n + 1
print(n)

print(name .. " " .. 10)  -- min 10

--comparison Operators
local age = 10

if age > 18 then 
  print("OVER 18")
end

-- else if 
age = 20

if age > 18 then
  print("over 18")
elseif age == 18 then
  print("18 huh")
else
  print("kiddo")
end


--bool
if flag then
  print("be grateful!")
end

if name ~= "min" then
  print("not min")
end



--combining statements
local age = 22
if age == 20 and x > 0 then 
  print("bravo")
elseif x == 18 or x > 18 then
  print("back up")
end
  
if not flag then
  print("false")
end

local function print_num(a)
  print(a)
end

local print_nums = function (a)
  print(a)
end

print_nums(5)
print_num(6)

function add(a, b)
  return a + b
end

add(2,43)

-- scope
function foo()
  local h = 10
end

print(h) -- nil

--loop 
--while 
local i = 1

while i < 10 do
  print("hi")
  i = i + 1
end

--for   
for i = 1, 10 do
  print("hello")
  i = i + 1
end


--arrays
local colors = {"red", "green", "blue"}
print(colors[1]

for i = 1, #colors do
  print(colors[i])
end

for index, value in ipairs(colors) do
 print(colors[index])
 
 print(value)
end

for _, value in ipairs(colors) do 
  print(value)
end

--Dictionaries
local info = {
  name = "min",
  age = 12
  flag = true
}

print(info["name"])
print(info.name)

for key, value in pairs(info) do
  print(key .. " " .. tostring(value))
end

--list
local data = {
  {"Sid", 20},
  {"mark", 23},
}

for i = 1, #data do 
 print(data[i][1] .. " is " .. data[i][2] .. "years old")
end

-- moudle
-- import code from one file
require("path")

require "custom"
