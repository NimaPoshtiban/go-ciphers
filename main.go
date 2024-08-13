package main

import (
	"context"
	"log"
	"net"

	"github.com/nimaposhtiban/go-ciphers/pb"
	"github.com/nimaposhtiban/go-ciphers/symmetric"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

type CipherGuideServer struct {
	pb.UnimplementedCipherServer
}

func (c *CipherGuideServer) proccess(ctx context.Context, in *pb.CipherRequest) (*pb.CipherResponse, error) {
	if in.Cipher {
		if in.Text == "" {
			return nil, proto.Error
		}
		switch in.Type {
		case pb.CipherRequest_CEASAR:
			var c symmetric.Caesar
			c = symmetric.Caesar(in.Text)
			result, _ := c.Encrypt()
			return &pb.CipherResponse{Text: result}, nil
		case pb.CipherRequest_VIGENERE:
			if in.Key == nil {
				return nil, proto.Error
			}
			var v symmetric.Vigenere
			v = symmetric.Vigenere(in.Text)
			result, _ := v.Encrypt(*in.Key)
			return &pb.CipherResponse{Text: result}, nil
		case pb.CipherRequest_OTP:
			if in.Key == nil {
				return nil, proto.Error
			}
			var o symmetric.OneTimePad
			o.Key = *in.Key
			o.PlainText = in.Text
			result := o.Encrypt()
			return &pb.CipherResponse{Text: result}, nil
		default:
			return nil, proto.Error
		}

	} else {
		if in.Text == "" {
			return nil, proto.Error
		}
		switch in.Type {
		case pb.CipherRequest_CEASAR:
			var c symmetric.Caesar
			c = symmetric.Caesar(in.Text)
			result, _ := c.Decrypt()
			return &pb.CipherResponse{Text: result}, nil
		case pb.CipherRequest_VIGENERE:
			if in.Key == nil {
				return nil, proto.Error
			}
			var v symmetric.Vigenere
			v = symmetric.Vigenere(in.Text)
			result, _ := v.Decrypt(*in.Key)
			return &pb.CipherResponse{Text: result}, nil
		case pb.CipherRequest_OTP:
			if in.Key == nil {
				return nil, proto.Error
			}
			var o symmetric.OneTimePad
			o.Key = *in.Key
			o.Hash = in.Text
			result := o.Decrypt()
			return &pb.CipherResponse{Text: result}, nil
		default:
			return nil, proto.Error
		}
	}
}
func newServer() *CipherGuideServer {
	s := &CipherGuideServer{}
	return s
}

func main() {
	li, err := net.Listen("tcp", "localhost:7042")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opt []grpc.ServerOption

	grpcServer := grpc.NewServer(opt...)
	pb.RegisterCipherServer(grpcServer, newServer())
	grpcServer.Serve(li)
}
