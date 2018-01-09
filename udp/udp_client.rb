
require 'socket'

udp_socket = UDPSocket.open
sockaddr = Socket.pack_sockaddr_in(8888, '127.0.0.1')
udp_socket.send('Hellow World!', 0, sockaddr)
udp_socket.close
