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
