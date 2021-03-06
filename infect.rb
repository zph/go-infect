module Infect
  VERSION = "0.0.5"
end

require 'open-uri'
require 'fileutils'

module Infect
  module Colorize
    def colorize(code, str)
      "\e[#{code}m#{str}\e[0m"
    end
    def notice(str)
      puts colorize(32, str)
    end
    def error(str)
      puts colorize(31, str)
    end
  end
end


require 'open-uri'
require 'fileutils'

module Infect
  class Command
    include Infect::Colorize

    def self.build(command, arg ,opts)
      case command.to_sym
      when :bundle
        Bundle.new(arg, opts)
      else
        $stderr.puts "WARNING: #{command} is not a valid command, ignorning"
      end
    end

    protected

    def mkdir(path)
      expanded_path = File.expand_path(path)
      unless File.directory?(expanded_path)
        notice "Making dir #{path}"
        FileUtils.mkdir_p(expanded_path)
      end
    end

    def chdir(path)
      Dir.chdir(path)
    end

    def download(url, path)
      File.open(File.expand_path(path), "w") do |file|
        open(url) do |read_file|
          file.write(read_file.read)
        end
      end
    end
  end
end

module Infect
  class Command
    class Bundle < Command
      attr_reader :bundle, :name, :location
      def initialize(arg, opts)
        @bundle = arg
        @options = opts
        @name = File.basename(bundle)
        @location = File.expand_path("#{BUNDLE_DIR}/#{name}")
      end

      def url
        "git@github.com:#{bundle}.git"
      end

      def install
        notice "Installing #{name}... "
        mkdir BUNDLE_DIR
        chdir BUNDLE_DIR
        git "clone '#{url}'"
      end

      def update
        notice "Updating #{name}... "
        chdir @location
        git "pull"
      end

      def call
        if File.exists? @location
          update
        else
          install
        end
      end

      private

      def git(args)
        `git #{args}`
      end
    end
  end
end

require 'fileutils'
module Infect
  class Command
    class Prereqs < Command
      def mkdirs(list)
        list.each do |path|
          FileUtils.mkdir_p(File.expand_path(path))
        end
      end
      def call
        mkdir "~/.vim/bundle"
        if RUBY_PLATFORM =~ /darwin/
          mkdirs %w(~/Library/Vim/swap ~/Library/Vim/backup ~/Library/Vim/undo)
        else
          mkdirs %w(~/.local/share/vim/swap ~/.local/share/vim/backup ~/.local/share/vim/undo")
        end
      end
    end
  end
end

module Infect
  class Cleanup
    include Infect::Colorize
    attr_reader :commands, :force

    def initialize(commands, args = {})
      @commands = commands
      @force = args[:force] || false
    end

    def call
      uninstall_unless_included names
    end

    private

    def uninstall_unless_included(list)
      Dir["#{BUNDLE_DIR}/*"].each do |path|
        unless list.include? File.basename(path)
          if confirm(path)
            notice "Deleting #{path}"
            require 'fileutils'
            FileUtils.rm_rf path
          else
            notice "Leaving #{path}"
          end
        end
      end
    end

    def confirm(name)
      unless force
        print "Remove #{name}? [Yn]: "
        response = STDIN.gets.chomp
        case response.downcase
        when ''
          true
        when 'y'
          true
        else
          false
        end
      end
    end

    def names
      list = []
      commands.each do |command|
        if command.respond_to? :name
          list << command.name
        end
      end
      list
    end

  end
end


module Infect
  VIMHOME = ENV['VIM'] || "#{ENV['HOME']}/.vim"
  VIMRC = ENV['MYVIMRC'] || "#{ENV['HOME']}/.vimrc"
  BUNDLE_DIR = "#{VIMHOME}/bundle"

  class Runner
    def self.call(*args)
      force = args.include? "-f"

      commands = [Command::Prereqs.new()]

      File.open( VIMRC ).each_line do |line|
        if line =~ /^"=/
          command, arg, opts = parse_command(line.gsub('"=', ''))
          commands << Command.build(command, arg, opts)
        end
      end

      commands.compact.peach(&:call)

      Cleanup.new(commands, :force => force).call

    end

    private

    def self.parse_command(line)



      command, arg, opts_string = line.split ' ', 3
      [command, arg, parse_opts(opts_string)]
    end

    def self.parse_opts(string)
      hash = {}
      parts = string.split(/[\s:](?=(?:[^"]|"[^"]*")*$)/).reject! { |c| c.empty? }
      if parts
        parts.each_slice(2) do |key, val|
          hash[key.to_sym] = val
        end
      end
      hash
    end
  end
end

Infect::Runner.call(*ARGV)
