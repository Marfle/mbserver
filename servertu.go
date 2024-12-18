package mbserver

import (
	"errors"
	"io"
	"log"
	"os"

	"github.com/goburrow/serial"
)

// ListenRTU starts the Modbus server listening to a serial device.
// For example:  err := s.ListenRTU(&serial.Config{Address: "/dev/ttyUSB0"})
func (s *Server) ListenRTU(serialConfig *serial.Config) (err error) {
	port, err := serial.Open(serialConfig)
	if err != nil {
		log.Fatalf("failed to open %s: %v\n", serialConfig.Address, err)
	}
	s.ports = append(s.ports, port)

	s.portsWG.Add(1)
	go func() {
		defer s.portsWG.Done()
		s.acceptSerialRequests(port)
	}()

	return err
}

func (s *Server) acceptSerialRequests(port serial.Port) {
SkipFrameError:
	for {
		select {
		case <-s.portsCloseChan:
			return
		default:
		}

		buffer := make([]byte, 512)
		haveBytes := 0

	readloop:
		for {
			bytesRead, err := port.Read(buffer[haveBytes:])
			if err != nil {
				if errors.Is(err, os.ErrDeadlineExceeded) {
					if haveBytes > 0 {
						log.Printf("timeout discarding buffered invalid data %v\n", buffer[:haveBytes])
					}
					continue SkipFrameError
				}
				if !errors.Is(err, io.EOF) {
					log.Printf("serial read error %v\n", err)
				}

				return
			}
			haveBytes += bytesRead

			log.Printf("serial read %v now %v\n", bytesRead, buffer[:haveBytes])

			if haveBytes >= 5 {

				// Set the length of the packet to the number of read bytes.
				packet := buffer[:haveBytes]

				// TODO don't complain on too short data until timeout or buffer full

				frame, err := NewRTUFrame(packet)
				if err != nil {
					log.Printf("serial frame warn %v\n", err)
					continue readloop
				}

				request := &Request{port, frame}

				log.Printf("correct request %+v\n", request.frame)

				s.requestChan <- request

				continue SkipFrameError
			}

			if haveBytes == 512 {
				log.Printf("receive discarding buffered invalid data %v\n", buffer)
				continue SkipFrameError
			}
		}
	}
}
