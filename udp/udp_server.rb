
require 'socket'

udp_socket = UDPSocket.open
udp_socket.bind('0.0.0.0', 8888)
p udp_socket.recv(65535)
udp_socket.close
