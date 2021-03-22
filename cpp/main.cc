// #define NDEBUG 1  

#include <unistd.h>
#include <stdio.h>
#include <gflags/gflags.h>
#include <glog/logging.h>


void quick_start() {
  LOG(INFO) << "Found " << 10 << " cookies";
  LOG(ERROR) << "this is serious";
}

void condition() {
  int count = 11;
  LOG_IF(INFO, count > 10) << "Got lots of items";
  LOG_IF(INFO, count > 100) << "Too many items";
}

void occasional() {
  for (int i = 1; i <= 20; i++) {
    LOG_EVERY_N(INFO, 10) << "Got the " << i << "th cookie";
  }

  for (int j = 1; j <= 10; j++) {
    LOG_FIRST_N(INFO, 5) << "log " << j << "th time";
  }
}

void debug_mode() {
  // If NDEBUG is defined, the following message will not be logged.
  DLOG(INFO) << "dlog info message";
}

void check() {
  CHECK_NE(1, 1) << ": The world must be ending";
  LOG(INFO) << "This mesage should not show up";
}

void vlog() {
  VLOG(1) << "I’m printed when you run the program with --v=1 or higher";
  VLOG(2) << "I’m printed when you run the program with --v=2 or higher";

  if (VLOG_IS_ON(2)) {
    printf("vlog 2 is on\n");
  }
}

int main(int argc, char* argv[]) {
  gflags::ParseCommandLineFlags(&argc, &argv, true);
  google::InitGoogleLogging(argv[0]);

  // quick_start();
  // condition();
  // occasional();
  // debug_mode();
  // check();
  vlog();
}
