# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: user.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("user.proto", :syntax => :proto3) do
    add_message "userpb.User" do
      optional :name, :string, 1
      optional :email, :string, 2
      optional :age, :int64, 3
      optional :type, :int64, 4
    end
  end
end

module Portfolio
  module UserPb
    User = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("userpb.User").msgclass
  end
end
