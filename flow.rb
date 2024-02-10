# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Flow < Formula
  desc ""
  homepage "https://github.com/Edens-Angel/flow-cli"
  version "0.1"

  on_macos do
    url "https://github.com/Edens-Angel/flow-cli/releases/download/v0.1/flow-cli_0.1_darwin_all.tar.gz"
    sha256 "eebd1aa999f391347b0aeebc198b7a429a9f3a0d79a3c8c9aea4f793ca2190be"

    def install
      bin.install "flow-cli"
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/Edens-Angel/flow-cli/releases/download/v0.1/flow-cli_0.1_linux_arm64.tar.gz"
      sha256 "93d86942065bf734e789337226971b3bef782dcca6250f6059ebaeadcc7e1e38"

      def install
        bin.install "flow"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/Edens-Angel/flow-cli/releases/download/v0.1/flow-cli_0.1_linux_amd64.tar.gz"
      sha256 "ba964329ef4557f1d9b05d31ada55500992b98d658fa1add1ad1f154ee1793ae"

      def install
        bin.install "flow"
      end
    end
  end
end
